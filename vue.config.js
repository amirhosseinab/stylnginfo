const path = require('path');

module.exports = {
    baseUrl: "dist",
    css: {
        loaderOptions: {
            sass: {
                data: '@import "assets/style/main";'
            }
        }
    },
    chainWebpack: config => {
        config
            .plugin('html')
            .tap(args => {
                args[0].template = path.resolve(__dirname, 'cmd/templates/index.gohtml');
                args[0].minify = false;
                args[0].hash = true;
                return args;
            })
    }
};