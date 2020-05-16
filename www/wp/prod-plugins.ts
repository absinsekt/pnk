import * as MiniCssExtractPlugin from 'mini-css-extract-plugin';
import * as HtmlWebpackPlugin from 'html-webpack-plugin';
import { CleanWebpackPlugin } from 'clean-webpack-plugin';
import { CleanCSSPlugin } from './clean-css-plugin';
import * as CopyWebpackPlugin from 'copy-webpack-plugin';
import { DefinePlugin } from 'webpack';
import { getEnvVariables } from './env-variables';

export const ProdPlugins = [
  new CleanWebpackPlugin(),

  new CopyWebpackPlugin([
    { from: 'src/assets/img', to: 'assets/img' },
  ]),

  new CleanCSSPlugin(),

  new MiniCssExtractPlugin({
    filename: '[name][hash].css'
  }),

  new HtmlWebpackPlugin({
    template: 'src/index.html',
    filename: '../../app/templates/_base.html'
  }),

  new DefinePlugin(getEnvVariables())
];
