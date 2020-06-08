const path = require('path');
const autoPreprocess = require('svelte-preprocess');
const webpack = require('webpack');

const {
  aliases,
  definedEnvVariables,
  printEnvVariables,
  ProdLoaders
} = require('../build-tools');

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

  const svelteLoader = config.module.rules.find(r => r.loader && r.loader.includes('svelte-loader'));
  svelteLoader.options = {
    ...svelteLoader.options,
    preprocess: [autoPreprocess({
      stylus: {
        paths: [path.resolve(__dirname, '../src')]
      }
    })]
  };

  config.resolve.alias = aliases;
  config.resolve.extensions = ['.mjs', '.ts', '.js', '.svelte', '.styl'];

  config.plugins.push(new webpack.DefinePlugin(definedEnvVariables));

  return config;
};
