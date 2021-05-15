"use strict";
const Generator = require("yeoman-generator");
const chalk = require("chalk");
const yosay = require("yosay");

module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);
    this.configKeys = ["appName", "appSecretKey", "jwtSecretKey", "serverPort", "websocket", "searchEngine", "messageBroker", "skipClient"];
  }

  prompting() {
    // Have Yeoman greet the user.
    this.log(
      yosay(`Welcome to the fine ${chalk.red("generator-gin")} generator!`)
    );

    const prompts = [];

    if (this.config.get("appName") === undefined) {
      prompts.push({
        type: "input",
        name: "appName",
        message: "Your project name?",
        default: this.appname
      })
    }
    if (this.config.get("jwtSecretKey") === undefined) {
      prompts.push({
        type: "input",
        name: "jwtSecretKey",
        message: "Your JWT secret key?",
        default: this._jwtGenKey()
      })
    }
    if (this.config.get("serverPort") === undefined) {
      prompts.push({
        type: "input",
        name: "serverPort",
        message: "Your server port?",
        default: "8080"
      })
    }
    if (this.config.get("appSecretKey") === undefined) {
      prompts.push({
        type: "input",
        name: "appSecretKey",
        message: "Your App secret key?",
        default: this._jwtGenKey()
      })
    }
    if (this.config.get("websocket") === undefined) {
      prompts.push({
        type: "confirm",
        name: "websocket",
        message: "Would you like to enable the web socket?",
        default: false
      })
    }
    if (this.config.get("searchEngine") === undefined) {
      prompts.push({
        type: "confirm",
        name: "searchEngine",
        message: "Would you like to enable the search engine?",
        default: false
      })
    }
    if (this.config.get("messageBroker") === undefined) {
      prompts.push({
        type: "confirm",
        name: "messageBroker",
        message: "Would you like to enable the message broker?",
        default: false
      })
    }
    if (this.config.get("skipClient") === undefined) {
      prompts.push({
        type: "confirm",
        name: "skipClient",
        message: "Would you like to skip client?",
        default: true
      })
    }

    return this.prompt(prompts).then(props => {
      this.props = props;
      this.props.destinationPath = this.destinationPath();
      
      this.configKeys.forEach(k => {
        if (this.config.get(k) !== undefined) {
          this.props[k] = this.config.get(k);
        }
      });
    });
  }

  end() {
    this.log(yosay(`Saving config!`));
    this.configKeys.forEach(k => {
      this.config.set(k, this.props[k]);
    });
    this.config.set("databaseType", "sql");
    this.config.set("authenticationType", "jwt");
    this.config.save();
  }

  _jwtGenKey() {
    return (
      Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15)
    );
  }

  writing() {
    this.fs.copyTpl(
      `${this.templatePath()}/**/!(_)*`,
      this.destinationPath(),
      this.props
    );
    this.fs.copyTpl(
      this.templatePath(".gitignore"),
      this.destinationPath(".gitignore"),
      this.props
    );
  }

  install() {
    this.spawnCommand("make", ["docs"]);
    this.spawnCommand("go", ["mod", "vendor"]);
  }
};
