import webpack from 'webpack';
import * as path from 'path';
import { autoPreprocess } from 'svelte-preprocess/dist/autoProcess';

import {
  aliases,
  DevLoaders,
  DevPlugins,
  ProdLoaders,
  ProdPlugins,
  printEnvVariables,
} from './build-tools';

const config: (env, argv) => webpack.Configuration = (env, argv) => {
  const isProduction = argv.mode === 'production';

  printEnvVariables();

  return {
    entry: {
      bundle: [
        './src/styles/index',
        './src/assets/icons.font',
        './src/app/index',
      ]
    },

    resolve: {
      extensions: ['.mjs', '.ts', '.js', '.svelte', '.styl'],
      modules: [
        'node_modules',
        './src/app',
      ],
      alias: aliases
    },

    output: {
      path: path.resolve(__dirname, 'dist'),
      filename: isProduction
        ? '[name][hash].js'
        : '[name].js',

      publicPath: '/dist/'
    },

    devServer: {
      contentBase: 'dist',
      public: '127.0.0.1:5001',
      port: 5001,
      disableHostCheck: true,
      hot: false
    },

    watch: !isProduction,

    module: {
      rules: [{
        test: /\.ts$/,
        include: [
          path.resolve(__dirname, 'src/app'),
          path.resolve(__dirname, 'src/pnk'),
        ],
        use: {
          loader: 'ts-loader',
        }
      }, {
        test: /\.svelte$/,
        use: {
          loader: 'svelte-loader',
          options: {
            emitCss: true,
            hotReload: false, //if true - Cannot read property '_debugName' of undefined
            preprocess: [autoPreprocess({
              stylus: {
                paths: [path.resolve(__dirname, 'src')]
              }
            })],
          },
        },
      }, {
        test: /\.css$/,
        use: isProduction ? ProdLoaders.cssLoaders : DevLoaders.cssLoaders
      }, {
        test: /\.styl$/,
        use: isProduction ? ProdLoaders.stylusLoaders : DevLoaders.stylusLoaders
      }, {
        test: /\.(svg|eot|woff|ttf|woff2|otf)$/,
        loader: {
          loader: 'file-loader',
          options: {
            name: 'assets/fonts/[name].[ext]',
            context: path.join(__dirname, 'src/')
          }
        }
      }, {
        test: /\.(png|jpg|gif)$/,
        loader: {
          loader: 'file-loader',
          options: {
            name: '[path][name].[ext]',
            context: path.join(__dirname, 'src/')
          }
        }
      }, {
        test: /\.font\.js$/,
        use: isProduction ?
          ProdLoaders.fontLoaders :
          DevLoaders.fontLoaders
      }]
    },

    mode: argv.mode,
    plugins: isProduction ? ProdPlugins : DevPlugins,
    devtool: isProduction ? false : 'source-map',

    optimization: {
      concatenateModules: true,
      removeAvailableModules: true,
      removeEmptyChunks: true,
      splitChunks: {
        cacheGroups: {
          default: false,
          vendor: {
            test: /node_modules/,
            name: 'vendor',
            chunks: 'all'
          }
        },
      },

      minimize: isProduction
    }
  };
}

export default config;
