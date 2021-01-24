"use strict";
const Generator = require("yeoman-generator");
const chalk = require("chalk");
const yosay = require("yosay");

module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);

    this.argument("entity", { type: String, required: true });

    if (this.options.help) return;

    // And you can then access it later; e.g.
    this.log(this.options.entity);
  }

  prompting() {
    const prompts = [
      {
        type: 'input',
        name: 'useRepoProxy',
        message: 'Use repository proxy ?',
        default: false
      },
      {
        type: 'input',
        name: 'useServiceProxy',
        message: 'Use service proxy ?',
        default: false
      }
    ];

    return this._optionOrPrompt(prompts).then(props => {
      this.props = props;
    });
  }

  start() {
    this.options.entityLower = this.options.entity.toLowerCase();
    this.options.entityCap = this._capitalize(this.options.entity);
    this.options.appName = this.config.get("appName");
  }

  _lower(word) {
    return word.charAt(0).toLowerCase() + word.slice(1);
  }

  _capitalize(word) {
    return word.charAt(0).toUpperCase() + word.slice(1);
  }

  writing() {
    const entityLower = this.options.entityLower;
    const entityCap = this.options.entityCap;

    this.fs.copyTpl(
      this.templatePath("controller/_temp.go"),
      this.destinationPath(`controller/${entityCap}.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("dto/_temp.go"),
      this.destinationPath(`dto/${entityCap}DTO.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("models/_temp.go"),
      this.destinationPath(`models/${entityCap}.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("repository/_temp.go"),
      this.destinationPath(`repository/${entityCap}.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("repository/impl/_temp.go"),
      this.destinationPath(`repository/impl/${entityCap}.go`),
      this.options
    );
    if (this.props.useRepoProxy === true) {
      this.fs.copyTpl(
        this.templatePath("repository/proxy/_temp.go"),
        this.destinationPath(`repository/proxy/${entityCap}.go`),
        this.options
      );
    }
    this.fs.copyTpl(
      this.templatePath("service/_temp.go"),
      this.destinationPath(`service/${entityCap}.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("service/impl/_temp.go"),
      this.destinationPath(`service/impl/${entityCap}.go`),
      this.options
    );
    if (this.props.useServiceProxy === true) {
      this.fs.copyTpl(
        this.templatePath("service/proxy/_temp.go"),
        this.destinationPath(`service/proxy/${entityCap}.go`),
        this.options
      );
    }
    this.fs.copyTpl(
      this.templatePath("service/mapper/_temp.go"),
      this.destinationPath(`service/mapper/${entityCap}.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("service/mapper/impl/_temp.go"),
      this.destinationPath(`service/mapper/impl/${entityCap}.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("dto/request/_entity/CreateRequestDTO.go"),
      this.destinationPath(`dto/request/${entityLower}/CreateRequestDTO.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("dto/request/_entity/UpdateRequestDTO.go"),
      this.destinationPath(`dto/request/${entityLower}/UpdateRequestDTO.go`),
      this.options
    );
    this.fs.copyTpl(
      this.templatePath("dto/response/_entity/ListResponseDTO.go"),
      this.destinationPath(`dto/response/${entityLower}/ListResponseDTO.go`),
      this.options
    );
    this._registerController();
    this._registerEntityDB();
    this._registerRoutesPrivate();
    this._registerRoutesPublic();
    // this._registerSecurity();
  }

  _registerController() {
    const entityLower = this.options.entityLower;
    const entityCap = this.options.entityCap;

    var path = this.destinationPath(`config/controller.go`);
    var file = this.fs.read(path);

    var controllerGlobalDeclare = `${entityCap}Controller controller.${entityCap}`;
    file = this._addControllerDeclareLine(file, '// Controllers globale declare end : dont remove', controllerGlobalDeclare)

    var mapperDeclare = `${entityLower}Mapper := mapper_impl.New${entityCap}()`;
    file = this._addControllerDeclareLine(file, '// Mappers declare end : dont remove', mapperDeclare)

    var repositoryDeclare = `${entityLower}Repo := repository_impl.New${entityCap}(db)`;
    file = this._addControllerDeclareLine(file, '// Repositories declare end : dont remove', repositoryDeclare)

    if (this.props.useRepoProxy === true) {
      var repositoryProxyDeclare = `${entityLower}RepoProxy := repository_proxy.New${entityCap}(db)`;
      file = this._addControllerDeclareLine(file, '// Proxy Repositories declare end : dont remove', repositoryDeclare)
    }

    var serviceDeclare = `${entityLower}Service := service_impl.New${entityCap}(${entityLower}Repo${this.props.useRepoProxy === true ? 'Proxy' : ''}, ${entityLower}Mapper)`;
    file = this._addControllerDeclareLine(file, '// Services declare end : dont remove', serviceDeclare)

    if (this.props.useServiceProxy === true) {
      var serviceProxyDeclare = `${entityLower}ServiceProxy := service_proxy.New${entityCap}(${entityLower}Service)`;
      file = this._addControllerDeclareLine(file, '// Proxy Services declare end : dont remove', serviceProxyDeclare)
    }

    var controllerDeclare = `${entityCap}Controller = controller.New${entityCap}(${entityLower}Service${this.props.useServiceProxy === true ? 'Proxy' : ''})`;
    file = this._addControllerDeclareLine(file, '// Controllers declare end : dont remove', controllerDeclare)

    this.fs.write(path, file);
  }

  _addControllerDeclareLine(file, insertKey, value, postFixValue = '\n  ') {
    if (file.indexOf(value) != -1) {
      return file
    }
    var position = file.indexOf(insertKey);
    if (position == -1) {
      return file
    }
    file = [file.slice(0, position), value+postFixValue, file.slice(position)].join('');
    return file;
  }

  _registerEntityDB() {
    const entityLower = this.options.entityLower;
    const entityCap = this.options.entityCap;

    var path = this.destinationPath(`config/database.go`);
    var file = this.fs.read(path);

    var modelDeclare = `&models.${entityCap}{},`;

    file = this._addControllerDeclareLine(file, '// Models declare end : dont remove', modelDeclare, '\n    ')

    this.fs.write(path, file);
  }

  _registerRoutesPrivate() {
    const entityLower = this.options.entityLower;
    const entityCap = this.options.entityCap;

    var path = this.destinationPath(`routers/api.private.go`);
    var file = this.fs.read(path);

    var apiDeclare = `{
    private${entityCap}Routes := privateRoutes.Group("/${entityLower}")
    private${entityCap}Routes.POST("", config.${entityCap}Controller.Create)
    private${entityCap}Routes.PUT("/:id", config.${entityCap}Controller.Update)
    private${entityCap}Routes.DELETE("/:id", config.${entityCap}Controller.Delete)
  }`;

    file = this._addControllerDeclareLine(file, '// Api declare end : dont remove', apiDeclare)

    this.fs.write(path, file);
  }

  _registerRoutesPublic() {
    const entityLower = this.options.entityLower;
    const entityCap = this.options.entityCap;

    var path = this.destinationPath(`routers/api.public.go`);
    var file = this.fs.read(path);

    var apiDeclare = `{
    public${entityCap}Routes := publicRoutes.Group("/${entityLower}")
    public${entityCap}Routes.GET("", config.${entityCap}Controller.GetAll)
    public${entityCap}Routes.GET("/:id", config.${entityCap}Controller.GetDetails)
  }`;

    file = this._addControllerDeclareLine(file, '// Api declare end : dont remove', apiDeclare)

    this.fs.write(path, file);
  }

  _registerSecurity() {
    const entityLower = this.options.entityLower;
    const entityCap = this.options.entityCap;

    var path = this.destinationPath(`middlewares/security.go`);
    var file = this.fs.read(path);

    var securityDeclare = `accessibleRoles["/api/private/${entityLower}.*"] = []string{constants.ROLE.ADMIN}`;

    file = this._addControllerDeclareLine(file, '// Security declare end : dont remove', securityDeclare)

    this.fs.write(path, file);
  }
};
