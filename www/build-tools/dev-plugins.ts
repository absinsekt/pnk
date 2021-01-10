import * as WriteFileWebpackPlugin from 'write-file-webpack-plugin';
import * as HtmlWebpackPlugin from 'html-webpack-plugin';
// import * as CopyWebpackPlugin from 'copy-webpack-plugin';
import { DefinePlugin } from 'webpack';
import { definedEnvVariables } from './env-variables';

export const DevPlugins = [
  new WriteFileWebpackPlugin(),

  // new CopyWebpackPlugin({
  //   patterns: [
  //     { from: 'src/assets/img', to: 'assets/img' },
  //   ],
  //   options: {
  //     concurrency: 100
  //   }
  // }),

  new HtmlWebpackPlugin({
    template: 'src/index.html',
    filename: '../../templates/shared/base.html'
  }),

  new DefinePlugin(definedEnvVariables)
];
