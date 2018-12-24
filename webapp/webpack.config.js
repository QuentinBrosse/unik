const path = require('path');
const fs = require('fs');

const certifsDir = path.join(__dirname, '../infra/data/certifs/local/');

module.exports = {
    mode: 'development',
    entry: {
        account: './account.js',
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ['babel-loader']
            }
        ]
    },
    resolve: {
        extensions: ['*', '.js', '.jsx']
    },
    devtool: 'inline-source-map',
    output: {
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, 'dist')
    },
    // watch: true,
    devServer: {
        contentBase: __dirname,
        publicPath: '/assets/',
        // compress: true,
        // port: 8000,
        https: {
            key: fs.readFileSync(path.join(certifsDir, 'localhost.key')),
            cert: fs.readFileSync(path.join(certifsDir, 'localhost.crt')),
            ca: fs.readFileSync(path.join(certifsDir, 'ca.pem')),
        }
    }
};
