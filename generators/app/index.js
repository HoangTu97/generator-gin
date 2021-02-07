"use strict";
const Generator = require("yeoman-generator");
const chalk = require("chalk");
const yosay = require("yosay");

module.exports = class extends Generator {
  prompting() {
    // Have Yeoman greet the user.
    this.log(
      yosay(`Welcome to the fine ${chalk.red("generator-gin")} generator!`)
    );

    const prompts = [
      {
        type: "input",
        name: "appName",
        message: "Your project name?",
        default: this.config.get("appName") || this.appname
      },
      {
        type: "input",
        name: "jwtSecretKey",
        message: "Your JWT secret key?",
        default: this.config.get("jwtSecretKey") || this._jwtGenKey()
      },
      {
        type: "input",
        name: "serverPort",
        message: "Your server port?",
        default: this.config.get("serverPort") || "8080"
      },
      {
        type: "input",
        name: "appSecretKey",
        message: "Your App secret key?",
        default: this.config.get("appSecretKey") || this._jwtGenKey()
      }
    ];

    return this.prompt(prompts).then(props => {
      this.props = props;
      this.props.destinationPath = this.destinationPath();
    });
  }

  _jwtGenKey() {
    return (
      Math.random()
        .toString(36)
        .substring(2, 15) +
      Math.random()
        .toString(36)
        .substring(2, 15)
    );
  }

  // _appGenKey() {
  //   var i,
  //     j,
  //     k = "";

  //   addEntropyTime();
  //   var seed = keyFromEntropy();

  //   var prng = new AESprng(seed);
  //   if (document.key.keytype[0].checked) {
  //     // Text key
  //     var charA = "A".charCodeAt(0);

  //     for (i = 0; i < 12; i++) {
  //       if (i > 0) {
  //         k += "-";
  //       }
  //       for (j = 0; j < 5; j++) {
  //         k += String.fromCharCode(charA + prng.nextInt(25));
  //       }
  //     }
  //   } else {
  //     // Hexadecimal key
  //     var hexDigits = "0123456789ABCDEF";

  //     for (i = 0; i < 64; i++) {
  //       k += hexDigits.charAt(prng.nextInt(15));
  //     }
  //   }
  //   return k;
  // }

  writing() {
    this.fs.copyTpl(
      `${this.templatePath()}/**/!(_)*`,
      this.destinationPath(),
      this.props
    );
  }

  install() {
    this.spawnCommand("make", ["docs"]);
    this.spawnCommand("go", ["mod", "vendor"]);
  }

  end() {
    this.log(yosay(`Saving config!`));
    this.config.set("appName", this.props.appName);
    this.config.set("appSecretKey", this.props.appSecretKey);
    this.config.set("jwtSecretKey", this.props.jwtSecretKey);
    this.config.set("serverPort", this.props.serverPort);
    this.config.save();
  }
};
