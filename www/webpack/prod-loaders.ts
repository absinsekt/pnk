import * as autoprefixer from 'autoprefixer';
import * as cssnano from 'cssnano';
import * as MiniCssExtractPlugin from 'mini-css-extract-plugin';

export const ProdLoaders = {
  cssLoaders: [
    MiniCssExtractPlugin.loader,
    {
      loader: 'postcss-loader',
      options: { plugins: [ autoprefixer(), cssnano() ] }
    },
    'css-loader',
  ],

  stylusLoaders: [
    MiniCssExtractPlugin.loader,
    'css-loader',
    {
      loader: 'postcss-loader',
      options: { plugins: [ autoprefixer(), cssnano() ] }
    },
    'stylus-loader'
  ],

  fontLoaders: [
    MiniCssExtractPlugin.loader,
    'css-loader',
    {
      loader: 'webfonts-loader',
      options: {
        embed: true
      }
    }
  ]
};
