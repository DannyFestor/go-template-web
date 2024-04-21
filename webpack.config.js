// webpack.config.js

const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = {
  mode: 'production',
  plugins: [new MiniCssExtractPlugin({ filename: 'app.css' })],
  entry: './resources/js/app.js',
  output: {
    path: path.resolve(__dirname, 'resources/public/assets/'),
    filename: 'app.js',
  },
  module: {
    rules: [
      {
        test: path.resolve(__dirname, 'resources/css/styles.css'),
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
            options: {
              emit: true,
            }
          },
          'css-loader',
          'postcss-loader'
        ],
      },
    ],
  },
};
