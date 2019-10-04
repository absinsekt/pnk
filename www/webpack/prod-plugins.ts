import * as MiniCssExtractPlugin from 'mini-css-extract-plugin';
import * as HtmlWebpackPlugin from 'html-webpack-plugin';
import { CleanWebpackPlugin } from 'clean-webpack-plugin';
import { CleanCSSPlugin } from './clean-css-plugin';

export const ProdPlugins = [
  new CleanWebpackPlugin(),

  new CleanCSSPlugin(),

  new MiniCssExtractPlugin({
    filename: '[name][hash].css'
  }),

  new HtmlWebpackPlugin({
    template: 'src/index.html',
    filename: '../../templates/base.html'
  })
];
