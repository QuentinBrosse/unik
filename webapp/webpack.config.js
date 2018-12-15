const path = require('path');
const fs = require('fs');

const certifsDir = path.join(__dirname, '../infra/data/certifs/local/');

module.exports = {
    mode: "development",
    entry: "./client.js",
    devtool: 'inline-source-map',
    output: {
        filename: 'app.bundle.js',
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
