import webpack from 'webpack';
import * as path from 'path';
import * as autoProcess from 'svelte-preprocess/src/autoProcess';
import * as stylus from 'svelte-preprocess/src/processors/stylus';

import {
  DevLoaders,
  DevPlugins,
  ProdLoaders,
  ProdPlugins,
} from './webpack';

const config: (env, argv) => webpack.Configuration = (env, argv) => {
  const isProduction = argv.mode === 'production';

  return {
    entry: {
      bundle: [
        './src/index',
        './src/styles/index'
      ]
    },

    resolve: {
      alias: {
        svelte: path.resolve('node_modules', 'svelte'),
        '@components': path.resolve(__dirname, 'src/app/components/'),
        '@assets': path.resolve(__dirname, 'src/assets/')
      },
      extensions: ['.ts', '.js', '.svelte', '.styl']
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
      public: '127.0.0.1',
      port: '5001',
      disableHostCheck: true,
      hot: false
    },

    watch: !isProduction,

    module: {
      rules: [{
        test: /\.svelte$/,
        exclude: /node_modules/,
        use: {
          loader: 'svelte-loader',
          options: {
            emitCss: true,
            hotReload: true,
            preprocess: [autoProcess(), stylus()]
          }
        }
      }, {
        test: /\.(j|t)s$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [
              ['@babel/preset-env', { modules: 'commonjs' }],
              '@babel/preset-typescript'
            ],
            plugins: [
              'transform-class-properties',
            ]
          }
        }
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
        test: /\.font\.js/,
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
