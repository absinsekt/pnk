import * as WriteFileWebpackPlugin from 'write-file-webpack-plugin';
import * as HtmlWebpackPlugin from 'html-webpack-plugin';
import * as CopyWebpackPlugin from 'copy-webpack-plugin';
import { DefinePlugin } from 'webpack';
import { definedEnvVariables } from './env-variables';

export const DevPlugins = [
  new WriteFileWebpackPlugin(),

  new CopyWebpackPlugin([
    { from: 'src/assets/img', to: 'assets/img' },
  ]),

  new HtmlWebpackPlugin({
    template: 'src/index.html',
    filename: '../../app/templates/_base.html'
  }),

  new DefinePlugin(definedEnvVariables)
];