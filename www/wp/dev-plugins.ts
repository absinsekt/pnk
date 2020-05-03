import * as WriteFileWebpackPlugin from 'write-file-webpack-plugin';
import * as HtmlWebpackPlugin from 'html-webpack-plugin';
import * as CopyWebpackPlugin from 'copy-webpack-plugin';

export const DevPlugins = [
  new WriteFileWebpackPlugin(),

  new CopyWebpackPlugin([
    { from: 'src/assets/img', to: 'assets/img' },
  ]),

  new HtmlWebpackPlugin({
    template: 'src/index.html',
    filename: '../../app/templates/_base.html'
  })
];
