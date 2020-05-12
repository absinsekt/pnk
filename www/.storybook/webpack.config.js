const path = require('path');
const autoPreprocess = require('svelte-preprocess');
const SRC_PATH = path.join(__dirname, '../src');
const STORIES_PATH = path.join(__dirname, '../stories');

module.exports = ({config}) => {
  config.module.rules.push({
    test: /\.[jt]s$/,
    include: [SRC_PATH, STORIES_PATH],
    use: [
      {
        loader: require.resolve('ts-loader')
      }
    ]
  });

  config.resolve.alias = {
    'app': path.resolve(__dirname, '../src/app'),
    'styles': path.resolve(__dirname, '../src/styles'),
    'assets': path.resolve(__dirname, '../src/assets'),
  };

  config.resolve.extensions = ['.mjs', '.ts', '.js', '.svelte', '.styl'];

  return config;
};
