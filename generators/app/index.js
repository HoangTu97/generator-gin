'use strict';
const Generator = require('yeoman-generator');
const chalk = require('chalk');
const yosay = require('yosay');

module.exports = class extends Generator {
  prompting() {
    // Have Yeoman greet the user.
    this.log(
      yosay(`Welcome to the fine ${chalk.red('generator-gin')} generator!`)
    );

    const prompts = [
      {
        type: 'input',
        name: 'appName',
        message: 'Your project name?',
        default: this.config.get("appName") || this.appname
      },
      {
        type: 'input',
        name: 'appSecretKey',
        message: 'Your JWT secret key?',
        default: this.config.get("jwtSecretKey") || Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15)
      },
      {
        type: 'input',
        name: 'serverPort',
        message: 'Your server port?',
        default: this.config.get("serverPort") || '8080'
      }
    ];

    return this.prompt(prompts).then(props => {
      this.props = props;
      this.props.destinationPath = this.destinationPath();
    });
  }

  writing() {
    this.fs.copyTpl(
      `${this.templatePath()}/**/!(_)*`,
      this.destinationPath(),
      this.props
    );
  }

  install() {
    this.spawnCommand('make', ['docs']);
    this.spawnCommand('go', ['mod','vendor']);
  }

  end() {
    this.log(
      yosay(`Saving config!`)
    );
    this.config.set("appName", this.props.appName)
    this.config.set("jwtSecretKey", this.props.appSecretKey)
    this.config.set("serverPort", this.props.serverPort)
    this.config.save();
  }
};
