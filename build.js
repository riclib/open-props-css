const fs = require('fs');
const path = require('path');
const postcss = require('postcss');
const postcssImport = require('postcss-import');

async function build() {
  const srcFile = path.join(__dirname, 'src', 'index.css');
  const outFile = path.join(__dirname, 'uicss', 'dashboard.css');

  // Ensure output directory exists
  if (!fs.existsSync(path.join(__dirname, 'uicss'))) {
    fs.mkdirSync(path.join(__dirname, 'uicss'));
  }

  try {
    const css = fs.readFileSync(srcFile, 'utf8');
    const result = await postcss([postcssImport])
      .process(css, { from: srcFile, to: outFile });

    fs.writeFileSync(outFile, result.css);
    console.log('✓ Built dashboard.css');
  } catch (err) {
    console.error('Build error:', err);
    process.exit(1);
  }
}

// Watch mode
if (process.argv.includes('--watch')) {
  console.log('Watching for changes...');
  fs.watch(path.join(__dirname, 'src'), { recursive: true }, (eventType, filename) => {
    if (filename && filename.endsWith('.css')) {
      console.log(`${filename} changed, rebuilding...`);
      build();
    }
  });
  build(); // Initial build
} else {
  build();
}