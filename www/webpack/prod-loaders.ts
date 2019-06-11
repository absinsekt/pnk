import * as MiniCssExtractPlugin from 'mini-css-extract-plugin';

export const ProdLoaders = {
  cssLoaders: [
    {
      loader: MiniCssExtractPlugin.loader,
    },
    'css-loader'
  ],

  stylusLoaders: [
    {
      loader: MiniCssExtractPlugin.loader,
    },
    'css-loader',
    {
      loader: 'postcss-loader',
      options: {
        sourceMap: true
      }
    },
    'stylus-loader'
  ],

  fontLoaders: [
    {
      loader: MiniCssExtractPlugin.loader,
    },
    'css-loader',
    {
      loader: 'webfonts-loader',
      options: {
        embed: true
      }
    }
  ]
};
