import autoprefixer from 'autoprefixer';
import * as cssnano from 'cssnano';

export const DevLoaders = {
  cssLoaders: [
    'style-loader',
    'css-loader'
  ],

  stylusLoaders: [
    'style-loader',
    'css-loader',
    {
      loader: 'postcss-loader',
      options: {
        postcssOptions: {
          plugins: [
            autoprefixer,
            cssnano
          ]
        },
        sourceMap: true
      }
    },
    'stylus-loader'
  ]
};
