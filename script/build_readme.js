const { join, resolve } = require('path');
const { readdirSync, statSync, readFileSync, writeFileSync } = require('fs');

const PROJECT_ROOT = join(resolve(__dirname, '..'));
const ALG_DIR = join(PROJECT_ROOT, 'algorithms');

const result = ['# LeetCode_Cs'];

const algNamePattern = /##\s\[(?<algName>.*)\]/;
const ALG_DATA = readdirSync(ALG_DIR)
    .filter((filename) => statSync(join(ALG_DIR, filename)).isDirectory())
    .map((filename) => {
        const firstLine = readFileSync(join(ALG_DIR, filename, 'README.md'))
            .toString()
            .split('\n')[0];

        const algName = firstLine.match(algNamePattern).groups.algName;
        if (algName === undefined) {
            throw new Error('can not find algName for ' + firstLine);
        }

        return {
            algName,
            filename,
            id: parseInt(algName.match(/^(?<id>\d+)/).groups.id),
        };
    })
    .sort((a, b) => a.id - b.id)
    .map(({ algName, filename }) => `[${algName}](algorithms/${filename})`)
    .join('\n\n');

result.push(
    `<details>\n<summary>Algorithms</summary>\n\n${ALG_DATA}\n\n</details>`
);
result.push('\n');

writeFileSync(join(PROJECT_ROOT, 'README.md'), result.join('\n\n'));
