const fs = require('fs');
const path = require('path');
const postcss = require('postcss');
const postcssImport = require('postcss-import');

async function build() {
  const srcFile = path.join(__dirname, 'src', 'index.css');
  const outFile = path.join(__dirname, 'uicss', 'dashboard.css');
  const readableFile = path.join(__dirname, 'uicss', 'dashboard.readable.css');

  // Ensure output directory exists
  if (!fs.existsSync(path.join(__dirname, 'uicss'))) {
    fs.mkdirSync(path.join(__dirname, 'uicss'));
  }

  try {
    const css = fs.readFileSync(srcFile, 'utf8');
    
    // Build minified version (normal)
    const result = await postcss([postcssImport])
      .process(css, { from: srcFile, to: outFile });

    fs.writeFileSync(outFile, result.css);
    console.log('✓ Built dashboard.css');

    // Build readable version with pretty formatting
    const readableResult = await postcss([
      postcssImport,
      {
        postcssPlugin: 'format-css',
        Once(root, { result }) {
          // Format the CSS for readability
          root.walkRules(rule => {
            // Add newline after opening brace
            rule.raws.after = '\n  ';
            // Add spaces around selectors
            rule.raws.between = ' ';
            // Add newline before closing brace
            rule.raws.semicolon = true;
          });
          
          root.walkDecls(decl => {
            // Add space after property colon
            decl.raws.between = ': ';
            // Ensure semicolons
            decl.raws.semicolon = true;
          });

          root.walkAtRules(atRule => {
            atRule.raws.after = '\n  ';
            atRule.raws.between = ' ';
          });

          root.walkComments(comment => {
            comment.raws.before = '\n\n';
            comment.raws.left = ' ';
            comment.raws.right = ' ';
          });
        }
      }
    ]).process(css, { from: srcFile, to: readableFile });

    // Add header comment to readable version
    const header = `/* 
 * Open Props CSS Framework - Readable Version
 * Generated: ${new Date().toISOString()}
 * 
 * This is a non-minified version for development and debugging.
 * Use dashboard.css for production.
 * 
 * Sections:
 * - Open Props Core
 * - Base Reset & Typography
 * - Layout Components
 * - Form Elements
 * - Buttons & Interactive
 * - Utilities
 * - Theme Support
 */

`;
    
    fs.writeFileSync(readableFile, header + readableResult.css);
    console.log('✓ Built dashboard.readable.css');
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