const { join, resolve } = require('path');
const { writeFileSync, mkdirSync } = require('fs');
const enquirer = require('enquirer');

const PROJECT_ROOT = join(resolve(__dirname, '..'));
const ALG_DIR = join(PROJECT_ROOT, 'algorithms');

const javascript = (problemDir) => {
    mkdirSync(join(problemDir, 'javascript'));
    writeFileSync(join(problemDir, 'javascript', 'index.js'), '');

    mkdirSync(join(problemDir, 'javascript', 'test'));
    writeFileSync(
        join(problemDir, 'javascript', 'test', 'index.js'),
        'const func = require(\'../index\');\nconst assert = require(\'power-assert\');\n'
    );
};

(async () => {
    const { problemName, problemLink, language } = await enquirer.prompt([
        {
            type: 'input',
            name: 'problemName',
            message: 'problem name',
        },
        {
            type: 'input',
            name: 'problemLink',
            message: 'problem link',
        },
        {
            type: 'select',
            name: 'language',
            message: 'which language do you use',
            choices: ['javascript'],
        },
    ]);

    const dirName = problemName.replace(/[.-\s]+/g, '-');

    mkdirSync(join(ALG_DIR, dirName));
    writeFileSync(
        join(ALG_DIR, dirName, 'README.md'),
        `## [${problemName}](${problemLink})\n\n### Description\n\n`
    );

    const mapping = { javascript };

    mapping[language](join(ALG_DIR, dirName));
})();
