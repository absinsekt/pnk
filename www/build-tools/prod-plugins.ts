import * as MiniCssExtractPlugin from 'mini-css-extract-plugin';
import * as HtmlWebpackPlugin from 'html-webpack-plugin';
import { CleanWebpackPlugin } from 'clean-webpack-plugin';
import { CleanCSSPlugin } from './clean-css-plugin';
// import * as CopyWebpackPlugin from 'copy-webpack-plugin';
import { DefinePlugin } from 'webpack';
import { definedEnvVariables } from './env-variables';

export const ProdPlugins = [
  new CleanWebpackPlugin(),

  // new CopyWebpackPlugin({
  //   patterns: [
  //     { from: 'src/assets/img', to: 'assets/img' },
  //   ],
  //   options: {
  //     concurrency: 100
  //   }
  // }),

  new CleanCSSPlugin(),

  new MiniCssExtractPlugin({
    filename: '[name][hash].css'
  }),

  new HtmlWebpackPlugin({
    template: 'src/index.html',
    filename: '../../templates/shared/base.html'
  }),

  new DefinePlugin(definedEnvVariables)
];
