import * as path from 'path';

import {
  DevLoaders,
  DevPlugins,
  ProdLoaders,
  ProdPlugins,
} from './webpack';

const config = (env, argv) => {
  const isProduction = argv.mode === 'production';
  const isDevelopment = argv.mode === 'development';
  const minimize = (isProduction || (env && env.mini === true));

  return {
    entry: {
      app: [
        'tsx/index',
        'styles/index',
        'vanilla/index'
      ]
    },

    resolve: {
      extensions: ['.ts', '.js', '.tsx', '.styl'],
      modules: [path.join(__dirname, 'www'), './node_modules'],
      alias: {
        assets: path.resolve(__dirname, 'www/assets')
      }
    },

    output: {
      path: path.join(__dirname, 'dist'),
      filename: isDevelopment
        ? '[name].js'
        : '[name][hash].js',

      publicPath: '/dist/'
    },

    devServer: {
      public: '127.0.0.1',
      port: '9001',
      disableHostCheck: true,
      hot: true
    },

    watch: isDevelopment,

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
      minimize
    },

    module: {
      rules: [{
        test: /\.(j|t)sx?$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
          // ,
          // options: {
          //   presets: [
          //     '@babel/preset-env'
          //   ]
          // }
        }
      }, {
        test: /\.css$/,
        use: isProduction ?
          ProdLoaders.cssLoaders :
          DevLoaders.cssLoaders
      }, {
        test: /\.styl$/,
        use: isProduction ?
          ProdLoaders.stylusLoaders :
          DevLoaders.stylusLoaders
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

    plugins: isProduction
      ? ProdPlugins
      : DevPlugins
  };
};

export default config;
