"use strict";
const Generator = require("yeoman-generator");
const chalk = require("chalk");
const yosay = require("yosay");
const _ = require('lodash');
var optionOrPrompt = require('yeoman-option-or-prompt');

module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);
    this._optionOrPrompt = optionOrPrompt;

    this.argument("schema", { type: String, required: true });

    if (this.options.help) return;

    var file = this.fs.readJSON(this.destinationPath(this.options.schema));
    this.options.config = file;
    this.options.useRepoProxy = this.options.config.useRepoProxy;
    this.options.useServiceProxy = this.options.config.useServiceProxy;
  }

  prompting() {
    const prompts = [];

    if (this.options.useRepoProxy === undefined) {
      prompts.push({
        type: 'input',
        name: 'useRepoProxy',
        message: 'Use repository proxy ?',
        default: false
      })
    }
    if (this.options.useServiceProxy === undefined) {
      prompts.push({
        type: 'input',
        name: 'useServiceProxy',
        message: 'Use service proxy ?',
        default: false
      })
    }

    return this._optionOrPrompt(prompts).then(props => {
      this.props = { ...this.props, props };
      if (this.props.useRepoProxy !== undefined) {
        this.options.useRepoProxy = this.props.useRepoProxy;
      }
      if (this.props.useServiceProxy !== undefined) {
        this.options.useServiceProxy = this.props.useServiceProxy;
      }
    });
  }

  start() {
    this.log(yosay("Analyzing model structure"));

    var modelRelationships = {};
    for (let index = 0; index < this.options.config.relationships.length; ++index) {
      var relationship = this.options.config.relationships[index];
      var from = relationship.from;
      var to = relationship.to;
      var type = relationship.type;
      if (modelRelationships[from] === undefined) {
        modelRelationships[from] = {};
      }
      modelRelationships[from][to] = {
        type: type
      };

      if (type === 'manyToMany') {
        if (modelRelationships[to] === undefined) {
          modelRelationships[to] = {};
        }
        var joinTable = _.snakeCase(from) + '_' + _.snakeCase(to) + 's';
        modelRelationships[to][from] = {
          type, joinTable
        };
        modelRelationships[from][to] = {
          type, joinTable
        };
      }
    }

    for (let index = 0; index < this.options.config.models.length; ++index) {
      var model = this.options.config.models[index];
      var entity = model.name;
      model.nameVar = _.camelCase(entity);
      model.nameClass = entity;
      model.relationships = modelRelationships[entity];

      for (let indexFi = 0; indexFi < model.fields.length; ++indexFi) {
        var field = model.fields[indexFi];

        var gorms = ''
        switch (field.type) {
          case "string":
            gorms += `type:varchar(255)`
            break;
          case "text":
            gorms += `type:text`
            field.type = 'string'
            break;
        }
        if (gorms !== '') {
          field.gormConfig = `gorm:"${gorms}"`
        } else {
          field.gormConfig = ''
        }

        var jsons = _.snakeCase(field.name);
        if (jsons !== '') {
          field.jsonConfig = `json:"${jsons}"`
        } else {
          field.jsonConfig = ''
        }
      }
    }
  }

  writing() {
    for (let index = 0; index < this.options.config.models.length; ++index) {
      var model = this.options.config.models[index];

      this._scafflodFiles(model.nameVar, model.nameClass, this.options.useRepoProxy, this.options.useServiceProxy, model.fields, model.relationships);
      this._registerController(model.nameVar, model.nameClass, this.options.useRepoProxy, this.options.useServiceProxy);
      this._registerEntityDB(model.nameVar, model.nameClass);
      this._registerRoutesPrivate(model.nameVar, model.nameClass);
      this._registerRoutesPublic(model.nameVar, model.nameClass);
      // this._registerSecurity(model.nameVar, model.nameClass);
    }
  }

  _scafflodFiles(entityVar, entityClass, useRepoProxy, useServiceProxy, fields, relationships) {
    var appName = this.config.get("appName");
    this.fs.copyTpl(
      this.templatePath("controller/_temp.go.ejs"),
      this.destinationPath(`controller/${entityClass}.go`),
      {entityVar, entityClass, appName, fields}
    );
    this.fs.copyTpl(
      this.templatePath("dto/_temp.go.ejs"),
      this.destinationPath(`dto/${entityClass}DTO.go`),
      {entityVar, entityClass, appName, fields}
    );
    this.fs.copyTpl(
      this.templatePath("models/_temp.go.ejs"),
      this.destinationPath(`models/${entityClass}.go`),
      {entityVar, entityClass, appName, fields, relationships}
    );
    this.fs.copyTpl(
      this.templatePath("repository/_temp.go.ejs"),
      this.destinationPath(`repository/${entityClass}.go`),
      {entityVar, entityClass, appName}
    );
    this.fs.copyTpl(
      this.templatePath("repository/impl/_temp.go.ejs"),
      this.destinationPath(`repository/impl/${entityClass}.go`),
      {entityVar, entityClass, appName}
    );
    if (useRepoProxy === true) {
      this.fs.copyTpl(
        this.templatePath("repository/proxy/_temp.go.ejs"),
        this.destinationPath(`repository/proxy/${entityClass}.go`),
        {entityVar, entityClass, appName}
      );
    }
    this.fs.copyTpl(
      this.templatePath("service/_temp.go.ejs"),
      this.destinationPath(`service/${entityClass}.go`),
      {entityVar, entityClass, appName}
    );
    this.fs.copyTpl(
      this.templatePath("service/impl/_temp.go.ejs"),
      this.destinationPath(`service/impl/${entityClass}.go`),
      {entityVar, entityClass, appName}
    );
    if (useServiceProxy) {
      this.fs.copyTpl(
        this.templatePath("service/proxy/_temp.go.ejs"),
        this.destinationPath(`service/proxy/${entityClass}.go`),
        {entityVar, entityClass, appName}
      );
    }
    this.fs.copyTpl(
      this.templatePath("service/mapper/_temp.go.ejs"),
      this.destinationPath(`service/mapper/${entityClass}.go`),
      {entityVar, entityClass, appName}
    );
    this.fs.copyTpl(
      this.templatePath("service/mapper/impl/_temp.go.ejs"),
      this.destinationPath(`service/mapper/impl/${entityClass}.go`),
      {entityVar, entityClass, appName, fields}
    );
    this.fs.copyTpl(
      this.templatePath("dto/request/_entity/CreateRequestDTO.go.ejs"),
      this.destinationPath(`dto/request/${entityVar}/CreateRequestDTO.go`),
      {entityVar, entityClass, appName, fields}
    );
    this.fs.copyTpl(
      this.templatePath("dto/request/_entity/UpdateRequestDTO.go.ejs"),
      this.destinationPath(`dto/request/${entityVar}/UpdateRequestDTO.go`),
      {entityVar, entityClass, appName, fields}
    );
    this.fs.copyTpl(
      this.templatePath("dto/response/_entity/ListResponseDTO.go.ejs"),
      this.destinationPath(`dto/response/${entityVar}/ListResponseDTO.go`),
      {entityVar, entityClass, appName, fields}
    );
  }

  _registerController(entityVar, entityClass, useRepoProxy, useServiceProxy) {
    var path = this.destinationPath(`config/controller.go`);
    var file = this.fs.read(path);

    var controllerGlobalDeclare = `${entityClass}Controller controller.${entityClass}`;
    file = this._insertLine(file, '// Controllers globale declare end : dont remove', controllerGlobalDeclare);

    var mapperDeclare = `${entityVar}Mapper := mapper_impl.New${entityClass}()`;
    file = this._insertLine(file, '// Mappers declare end : dont remove', mapperDeclare)

    var repositoryDeclare = `${entityVar}Repo := repository_impl.New${entityClass}(db)`;
    file = this._insertLine(file, '// Repositories declare end : dont remove', repositoryDeclare)

    if (useRepoProxy === true) {
      var repositoryProxyDeclare = `${entityVar}RepoProxy := repository_proxy.New${entityClass}(db)`;
      file = this._insertLine(file, '// Proxy Repositories declare end : dont remove', repositoryProxyDeclare)
    }

    var serviceDeclare;
    if (useRepoProxy === true) {
      serviceDeclare = `${entityVar}Service := service_impl.New${entityClass}(${entityVar}RepoProxy, ${entityVar}Mapper)`;
    } else {
      serviceDeclare = `${entityVar}Service := service_impl.New${entityClass}(${entityVar}Repo, ${entityVar}Mapper)`;
    }
    file = this._insertLine(file, '// Services declare end : dont remove', serviceDeclare)

    if (useServiceProxy === true) {
      var serviceProxyDeclare = `${entityVar}ServiceProxy := service_proxy.New${entityClass}(${entityVar}Service)`;
      file = this._insertLine(file, '// Proxy Services declare end : dont remove', serviceProxyDeclare)
    }

    var controllerDeclare;
    if (useServiceProxy === true) {
      controllerDeclare= `${entityClass}Controller = controller.New${entityClass}(${entityVar}ServiceProxy)`;
    } else {
      controllerDeclare= `${entityClass}Controller = controller.New${entityClass}(${entityVar}Service)`;
    }
    file = this._insertLine(file, '// Controllers declare end : dont remove', controllerDeclare)

    this.fs.write(path, file);
  }

  _insertLine(file, insertKey, value, postFixValue = '\n  ') {
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

  _registerEntityDB(entityVar, entityClass) {
    var path = this.destinationPath(`config/database.go`);
    var file = this.fs.read(path);

    var modelDeclare = `&models.${entityClass}{},`;

    file = this._insertLine(file, '// Models declare end : dont remove', modelDeclare, '\n    ')

    this.fs.write(path, file);
  }

  _registerRoutesPrivate(entityVar, entityClass) {
    var path = this.destinationPath(`routers/api.private.go`);
    var file = this.fs.read(path);

    var apiDeclare = `{
    private${entityClass}Routes := privateRoutes.Group("/${entityVar}")
    private${entityClass}Routes.POST("", config.${entityClass}Controller.Create)
    private${entityClass}Routes.PUT("/:id", config.${entityClass}Controller.Update)
    private${entityClass}Routes.DELETE("/:id", config.${entityClass}Controller.Delete)
  }`;

    file = this._insertLine(file, '// Api declare end : dont remove', apiDeclare)

    this.fs.write(path, file);
  }

  _registerRoutesPublic(entityVar, entityClass) {
    var path = this.destinationPath(`routers/api.public.go`);
    var file = this.fs.read(path);

    var apiDeclare = `{
    public${entityClass}Routes := publicRoutes.Group("/${entityVar}")
    public${entityClass}Routes.GET("", config.${entityClass}Controller.GetAll)
    public${entityClass}Routes.GET("/:id", config.${entityClass}Controller.GetDetails)
  }`;

    file = this._insertLine(file, '// Api declare end : dont remove', apiDeclare)

    this.fs.write(path, file);
  }

  _registerSecurity(entityVar, entityClass) {
    var path = this.destinationPath(`middlewares/security.go`);
    var file = this.fs.read(path);

    var securityDeclare = `accessibleRoles["/api/private/${entityVar}.*"] = []string{constants.ROLE.ADMIN}`;

    file = this._insertLine(file, '// Security declare end : dont remove', securityDeclare)

    this.fs.write(path, file);
  }
};
