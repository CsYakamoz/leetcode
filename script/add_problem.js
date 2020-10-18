const { join, resolve } = require('path');
const { writeFileSync, mkdirSync } = require('fs');
const { prompt } = require('enquirer');

const crawlProblemInfo = require('./crawl_problem_info');

const PROJECT_ROOT = join(resolve(__dirname, '..'));
const ALG_DIR = join(PROJECT_ROOT, 'algorithms');
const LEETCODE_SUPPORTED_LANGUAGE = [
    { name: 'javascript', isFamiliar: true },
    { name: 'cpp', isFamiliar: false },
    { name: 'java', isFamiliar: false },
    { name: 'python', isFamiliar: false },
    { name: 'python3', isFamiliar: false },
    { name: 'c', isFamiliar: false },
    { name: 'csharp', isFamiliar: false },
    { name: 'ruby', isFamiliar: false },
    { name: 'swift', isFamiliar: false },
    { name: 'golang', isFamiliar: false },
    { name: 'scala', isFamiliar: false },
    { name: 'kotlin', isFamiliar: false },
    { name: 'rust', isFamiliar: false },
    { name: 'php', isFamiliar: false },
    { name: 'typescript', isFamiliar: false },
];

// call the main function on the last line
const main = async () => {
    const { problemLink, language } = await prompt([
        {
            type: 'input',
            name: 'problemLink',
            message: 'problem link',
        },
        {
            type: 'select',
            name: 'language',
            message: 'which language do you use',
            choices: LEETCODE_SUPPORTED_LANGUAGE.filter(
                (item) => item.isFamiliar
            ).map((item) => item.name),
        },
    ]);

    const {
        problemName,
        description,
        tags,
        code,
        difficulty,
    } = await crawlProblemInfo(problemLink, language);

    const dirName = problemName
        .replace(/[.-\s]+/g, '-')
        .replace(/^(\d+)/, (match) => match.padStart(4, '0'));

    const problemDir = join(ALG_DIR, dirName);
    mkdirSync(problemDir);

    generateREADME(
        problemDir,
        problemName,
        problemLink,
        description,
        difficulty,
        tags
    );

    const mapping = { javascript };
    mapping[language](problemDir, problemName, code);
};

const generateREADME = (
    problemDir,
    problemName,
    problemLink,
    description,
    difficulty,
    tags
) => {
    writeFileSync(
        join(problemDir, 'README.md'),
        [
            `## [${problemName}](${problemLink})`,
            '',
            '### Description',
            '',
            description,
            '',
            `**Difficulty:** \`${difficulty}\``,
            '',
            `**Tags:** ${tags.map((item) => `\`${item}\``).join(' ')}`,
            '',
        ].join('\n')
    );
};

const javascript = (problemDir, problemName, code) => {
    mkdirSync(join(problemDir, 'javascript'));
    writeFileSync(join(problemDir, 'javascript', 'index.js'), code);

    mkdirSync(join(problemDir, 'javascript', 'test'));
    writeFileSync(
        join(problemDir, 'javascript', 'test', 'index.js'),
        [
            'const func = require(\'../index\');',
            'const assert = require(\'power-assert\');',
            '',
            `describe('${problemName}', () => {});`,
            '',
        ].join('\n')
    );
};

main().catch((e) => console.log(e.stack));
