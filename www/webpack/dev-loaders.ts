export const DevLoaders = {
  cssLoaders: [
    'style-loader',
    {
      loader: 'postcss-loader',
      options: { sourceMap: true }
    },
    'css-loader'
  ],

  stylusLoaders: [
    'style-loader',
    {
      loader: 'css-loader'
    },
    {
      loader: 'postcss-loader',
      options: {
        sourceMap: true
      }
    },
    'stylus-loader'
  ],

  fontLoaders: [
    'style-loader',
    'css-loader',
    {
      loader: 'webfonts-loader',
      options: { embed: true }
    }
  ]
};
