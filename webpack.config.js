const path = require("path");
const webpack = require("webpack");
const BundleTracker = require("webpack-bundle-tracker");

module.exports = {
  context: __dirname,
  entry: "./assets/js/index",
  module: {
    rules: [
      {
        test: /\.s[ac]ss$/i,
        use: ["style-loader", "css-loader", "sass-loader"],
      },
    ],
  },
  output: {
    path: path.resolve(__dirname, "dist/webpack_bundles/"),
    publicPath: "auto",
    filename: "[name]-[contenthash].js",
  },
  plugins: [
    new BundleTracker({ path: __dirname, filename: "webpack-stats.json" }),
    new webpack.ProvidePlugin({
      $: "jquery",
      jQuery: "jquery",
    }),
  ],
};
