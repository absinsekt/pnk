const path = require('path');
const autoPreprocess = require('svelte-preprocess');
const webpack = require('webpack');

const aliases = require('../wp/aliases').aliases;
const { getEnvVariables, printEnvVariables } = require('../wp/env-variables');

const SRC_PATH = path.join(__dirname, '../src');
const STORIES_PATH = path.join(__dirname, '../stories');

module.exports = ({config}) => {
  printEnvVariables();

  config.module.rules.push({
    test: /\.[jt]s$/,
    include: [SRC_PATH, STORIES_PATH],
    use: [
      {
        loader: require.resolve('ts-loader')
      }
    ]
  });

  config.resolve.alias = aliases;
  config.resolve.extensions = ['.mjs', '.ts', '.js', '.svelte', '.styl'];

  config.plugins.push(new webpack.DefinePlugin(getEnvVariables()));

  return config;
};
