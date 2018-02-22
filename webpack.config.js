'use strict';

var ExtractTextPlugin = require("extract-text-webpack-plugin")
var path = require('path');
const goStaticPath = './app/javascript/'
const goStyles = './app/styles/'

module.exports = {
  cache: true,
  entry: {
    app: goStaticPath + "base/app.js",
    user_register: goStaticPath + "users/register/register.js",
    user_login: goStaticPath + "users/login/login.js",
    user_logout: goStaticPath + "users/logout/logout.js",
    tokens_list: goStaticPath + "tokens/list/list.js",
    tokens_dapp: goStaticPath + "tokens/dapp/dapp.js",
    tokens_create: goStaticPath + "tokens/create/create.js",
    wallet_index: goStaticPath + "wallets/index/index.js",
    bundle: goStyles + "base.scss",
  },

  output: {
    path: path.resolve(__dirname, "app/assets"),
    filename: "javascript/[name].js"
  },

  module: {
    loaders: [
      {
        test: /\.css$/,
        loaders: ['style-loader','css-loader']
      },
      {
        test: /\.(scss|sass)$/,
        loader: ExtractTextPlugin.extract(["css-loader", "sass-loader"])
      },
      {
        test: /\.(png|jpg|jpeg|gif|svgz)(\?.+)?$/,
        use: [{
          loader: 'url-loader',
          options: {
            limit: 10000
          }
        }]
      },
      {
        test: /\.(woff|woff2)(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'url-loader?limit=10000&mimetype=application/font-woff'
      },
      {
         test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
         loader: 'url-loader?limit=10000&mimetype=application/octet-stream'
      },
      {
         test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
         loader: 'file-loader'
      },
      {
        test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'url-loader?limit=10000&mimetype=image/svg+xml'
      },
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
        query: {
            presets: ['es2015']
        }
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader'
      },
      {
        test: /\.js$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
          presets: ["es2015"]
        }
      },
    ]
  },
  plugins: [
    new ExtractTextPlugin("css/[name].css")
  ],
  node: {
    __dirname: true
  },
  resolve: {
    alias: {
      Chartist: "chartist/dist/chartist.js",
      jQuery: "jquery/dist/jquery.js",
      $: "jquery/dist/jquery.js",
      "window.jQuery":"jquery",
      Tether: 'tether',
      vue: 'vue/dist/vue.js',
      jsappbase: path.resolve(__dirname, "./app/javascript/base"),
    },
  },
}
