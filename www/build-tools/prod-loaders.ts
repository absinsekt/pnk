import autoprefixer from 'autoprefixer';
import * as cssnano from 'cssnano';

import * as MiniCssExtractPlugin from 'mini-css-extract-plugin';

export const ProdLoaders = {
  cssLoaders: [
    MiniCssExtractPlugin.loader,
    'css-loader',
  ],

  stylusLoaders: [
    MiniCssExtractPlugin.loader,
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
