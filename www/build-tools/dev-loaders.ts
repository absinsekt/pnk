import * as autoprefixer from 'autoprefixer';
import * as cssnano from 'cssnano';

export const DevLoaders = {
  cssLoaders: [
    'style-loader',
    'css-loader'
  ],

  stylusLoaders: [
    'style-loader',
    {
      loader: 'css-loader'
    },
    {
      loader: 'postcss-loader',
      options: { sourceMap: true, plugins: [ autoprefixer(), cssnano() ] }
    },
    'stylus-loader'
  ]
};
