const webpack = require('webpack');
const path = require('path');

const BUILD_DIR = path.resolve(__dirname, './public');
const APP_DIR = path.resolve(__dirname, './src');

const config = {
	mode: 'development',
	entry: {
		main: APP_DIR + '/index.js'
	},
	output: {
		filename: 'bundle.js',
		path: BUILD_DIR
	},
	module: {
		rules: [
			{
				test: /\.css$/,
				use: [{
					loader: 'style-loader'
				}, {
					loader: 'css-loader'
				}]
			}, {
				test: /\.(js|jsx)?$/,
				use: [{
					loader: 'babel-loader',
					options: {
						cacheDirectory: true,
						presets: ['@babel/react', '@babel/env']
					}
				}]
			}, {
				test: /\.svg$/,
				use: [{
					loader: 'svg-loader'
				}]
			}
		]
	}
};//-- end config

module.exports = config;

