{
  "env": {
      "browser": true,
      "commonjs": true,
      "es6": true,
      "node": true
  },
  "extends": [
      "eslint:recommended",
      "plugin:react/recommended"
  ],
  "parser": "babel-eslint",
  "parserOptions": {
      "ecmaFeatures": {
          "experimentalObjectRestSpread": true,
          "jsx": true,
          "legacyDecorators": true
      },
      "sourceType": "module",
      "allowImportExportEverywhere": false
  },
  "plugins": [
      "react"
  ],
  "rules": {
      "indent": [
          "off",
          4
      ],
      "linebreak-style": [
          "warn",
          "unix"
      ],
      "no-console": "off",
      "no-mixed-spaces-and-tabs": "warn",
      "no-unused-vars": "off",
      "react/prop-types": ["warn", {"skipUndeclared": true}],
      "quotes": [
          "off",
          "single"
      ],
      "semi": [
          "off",
          "always"
      ]
  }
}