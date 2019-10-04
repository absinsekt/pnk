import * as WriteFileWebpackPlugin from 'write-file-webpack-plugin';
import * as HtmlWebpackPlugin from 'html-webpack-plugin';

export const DevPlugins = [
  new WriteFileWebpackPlugin(),
  new HtmlWebpackPlugin({
    template: 'src/index.html',
    filename: '../../templates/base.html'
  })
];
