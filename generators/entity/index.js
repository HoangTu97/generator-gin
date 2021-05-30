"use strict";
const Generator = require("yeoman-generator");
// Const chalk = require("chalk");
// const yosay = require("yosay");
const _ = require("lodash");
var optionOrPrompt = require("yeoman-option-or-prompt");

module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);
    this._optionOrPrompt = optionOrPrompt;

    this.argument("entity", { type: String, required: true });

    if (this.options.help) return;

    // And you can then access it later; e.g.
    this.log(this.options.entity);
  }

  prompting() {
    const prompts = [];

    if (this.options.useRepoProxy === undefined) {
      prompts.push({
        type: "input",
        name: "useRepoProxy",
        message: "Use repository proxy ?",
        default: false
      });
    }

    if (this.options.useServiceProxy === undefined) {
      prompts.push({
        type: "input",
        name: "useServiceProxy",
        message: "Use service proxy ?",
        default: false
      });
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
    var entity = this.options.entity;
    var model = {};
    model.nameVar = _.camelCase(entity);
    model.nameClass = entity;
    this.options.model = model;
  }

  writing() {
    var model = this.options.model;
    this._scafflodFiles(
      model.nameVar,
      model.nameClass,
      this.options.useRepoProxy,
      this.options.useServiceProxy,
      [],
      []
    );
    this._registerController(
      model.nameVar,
      model.nameClass,
      this.options.useRepoProxy,
      this.options.useServiceProxy
    );
    this._registerEntityDB(model.nameVar, model.nameClass);
  }

  _scafflodFiles(
    entityVar,
    entityClass,
    useRepoProxy,
    useServiceProxy,
    fields,
    relationships
  ) {
    var appName = this.config.get("appName");
    this.fs.copyTpl(
      this.templatePath("controller/_temp.go.ejs"),
      this.destinationPath(`controller/${entityClass}.go`),
      { entityVar, entityClass, appName, fields }
    );
    this.fs.copyTpl(
      this.templatePath("dto/_temp.go.ejs"),
      this.destinationPath(`dto/${entityClass}DTO.go`),
      { entityVar, entityClass, appName, fields }
    );
    this.fs.copyTpl(
      this.templatePath("models/_temp.go.ejs"),
      this.destinationPath(`models/${entityClass}.go`),
      { entityVar, entityClass, appName, fields, relationships }
    );
    this.fs.copyTpl(
      this.templatePath("repository/_temp.go.ejs"),
      this.destinationPath(`repository/${entityClass}.go`),
      { entityVar, entityClass, appName }
    );
    this.fs.copyTpl(
      this.templatePath("repository/impl/_temp.go.ejs"),
      this.destinationPath(`repository/impl/${entityClass}.go`),
      { entityVar, entityClass, appName }
    );
    if (useRepoProxy === true) {
      this.fs.copyTpl(
        this.templatePath("repository/proxy/_temp.go.ejs"),
        this.destinationPath(`repository/proxy/${entityClass}.go`),
        { entityVar, entityClass, appName }
      );
    }

    this.fs.copyTpl(
      this.templatePath("service/_temp.go.ejs"),
      this.destinationPath(`service/${entityClass}.go`),
      { entityVar, entityClass, appName }
    );
    this.fs.copyTpl(
      this.templatePath("service/impl/_temp.go.ejs"),
      this.destinationPath(`service/impl/${entityClass}.go`),
      { entityVar, entityClass, appName }
    );
    if (useServiceProxy) {
      this.fs.copyTpl(
        this.templatePath("service/proxy/_temp.go.ejs"),
        this.destinationPath(`service/proxy/${entityClass}.go`),
        { entityVar, entityClass, appName }
      );
    }

    this.fs.copyTpl(
      this.templatePath("service/mapper/_temp.go.ejs"),
      this.destinationPath(`service/mapper/${entityClass}.go`),
      { entityVar, entityClass, appName }
    );
    this.fs.copyTpl(
      this.templatePath("service/mapper/impl/_temp.go.ejs"),
      this.destinationPath(`service/mapper/impl/${entityClass}.go`),
      { entityVar, entityClass, appName, fields }
    );
    this.fs.copyTpl(
      this.templatePath("dto/request/_entity/CreateRequestDTO.go.ejs"),
      this.destinationPath(`dto/request/${entityVar}/CreateRequestDTO.go`),
      { entityVar, entityClass, appName, fields }
    );
    this.fs.copyTpl(
      this.templatePath("dto/request/_entity/UpdateRequestDTO.go.ejs"),
      this.destinationPath(`dto/request/${entityVar}/UpdateRequestDTO.go`),
      { entityVar, entityClass, appName, fields }
    );
    this.fs.copyTpl(
      this.templatePath("dto/response/_entity/ListResponseDTO.go.ejs"),
      this.destinationPath(`dto/response/${entityVar}/ListResponseDTO.go`),
      { entityVar, entityClass, appName, fields }
    );
  }

  _registerController(entityVar, entityClass, useRepoProxy, useServiceProxy) {
    var path = this.destinationPath(`config/providers.go`);
    var file = this.fs.read(path);

    var mapperDeclare = `${entityVar}Mapper := mapper_impl.New${entityClass}()`;
    file = this._insertLine(
      file,
      "// Mappers declare end : dont remove",
      mapperDeclare
    );

    var repositoryDeclare = `${entityVar}Repo := repository_impl.New${entityClass}(db)`;
    file = this._insertLine(
      file,
      "// Repositories declare end : dont remove",
      repositoryDeclare
    );

    if (useRepoProxy === true) {
      var repositoryProxyDeclare = `${entityVar}RepoProxy := repository_proxy.New${entityClass}(db)`;
      file = this._insertLine(
        file,
        "// Proxy Repositories declare end : dont remove",
        repositoryProxyDeclare
      );
    }

    var serviceDeclare;
    if (useRepoProxy === true) {
      serviceDeclare = `${entityVar}Service := service_impl.New${entityClass}(${entityVar}RepoProxy, ${entityVar}Mapper)`;
    } else {
      serviceDeclare = `${entityVar}Service := service_impl.New${entityClass}(${entityVar}Repo, ${entityVar}Mapper)`;
    }

    file = this._insertLine(
      file,
      "// Services declare end : dont remove",
      serviceDeclare
    );

    if (useServiceProxy === true) {
      var serviceProxyDeclare = `${entityVar}ServiceProxy := service_proxy.New${entityClass}(${entityVar}Service)`;
      file = this._insertLine(
        file,
        "// Proxy Services declare end : dont remove",
        serviceProxyDeclare
      );
    }

    var controllerDeclare;
    if (useServiceProxy === true) {
      controllerDeclare = `${entityVar}Controller := controller.New${entityClass}(${entityVar}ServiceProxy)`;
    } else {
      controllerDeclare = `${entityVar}Controller := controller.New${entityClass}(${entityVar}Service)`;
    }

    file = this._insertLine(
      file,
      "// Controllers declare end : dont remove",
      controllerDeclare
    );

    var controllerGlobalDeclare = `${entityVar}Controller,`;
    file = this._insertLine(
      file,
      "// Register controller declare end : dont remove",
      controllerGlobalDeclare,
      "\n    "
    );

    this.fs.write(path, file);
  }

  _insertLine(file, insertKey, value, postFixValue = "\n  ") {
    if (file.indexOf(value) !== -1) {
      return file;
    }

    var position = file.indexOf(insertKey);
    if (position === -1) {
      return file;
    }

    file = [
      file.slice(0, position),
      value + postFixValue,
      file.slice(position)
    ].join("");
    return file;
  }

  _registerEntityDB(entityVar, entityClass) {
    var path = this.destinationPath(`config/database.go`);
    var file = this.fs.read(path);

    var modelDeclare = `&models.${entityClass}{},`;

    file = this._insertLine(
      file,
      "// Models declare end : dont remove",
      modelDeclare,
      "\n    "
    );

    this.fs.write(path, file);
  }
};
