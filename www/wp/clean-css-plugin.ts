import * as RawSource from 'webpack-sources/lib/RawSource';
import * as CleanCSS from 'clean-css';

export class CleanCSSPlugin {
  apply(compiler) {
    compiler.plugin('emit', function (compilation, next) {
      for (let fn in compilation.assets) {
        if (/(vendor|bundle)[a-z0-9]+\.css/.test(fn)) {
          let input = compilation.assets[fn].source(),
            output = new CleanCSS({ level: 2, compatibility: 'ie7'})
              .minify(input.replace(/\/\*[^*]*\*+([^/*][^*]*\*+)*\//gim, '')).styles;

          compilation.assets[fn] = new RawSource(output);
        }
      }

      next();
    });
  }
}
