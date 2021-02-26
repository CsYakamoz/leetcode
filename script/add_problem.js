const { join, resolve, relative } = require('path');
const { writeFileSync, mkdirSync, existsSync, statSync } = require('fs');
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
    { name: 'golang', isFamiliar: true },
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
            type: 'autocomplete',
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
    mkdir(problemDir);

    await generateREADME(
        problemDir,
        problemName,
        problemLink,
        description,
        difficulty,
        tags
    );

    const mapping = {
        javascript,
        golang,
    };
    await mapping[language](problemDir, problemName, code);
};

/**
 * @param {string} path
 */
const mkdir = (path) => {
    if (existsSync(path)) {
        if (!statSync(path).isDirectory()) {
            throw new Error(
                `path(${path}) already exists, but not a directory`
            );
        }
        return;
    }

    mkdirSync(path);
};

/**
 * @param {string} path
 * @param {string} content
 */
const writeFile = async (path, content) => {
    if (existsSync(path)) {
        if (!statSync(path).isFile()) {
            throw new Error(`path(${path}) already exists, but not a file`);
        }

        const relativePath = relative(PROJECT_ROOT, path);
        const { confirm } = await prompt([
            {
                type: 'confirm',
                name: 'confirm',
                message: `file(${relativePath}) already exists, overwrite?`,
            },
        ]);
        if (!confirm) {
            return;
        }
    }

    writeFileSync(path, content);
};

/**
 * @param {string} problemDir
 * @param {string} problemName
 * @param {string} problemLink
 * @param {string} description
 * @param {string} difficulty
 * @param {string[]} tags
 */
const generateREADME = async (
    problemDir,
    problemName,
    problemLink,
    description,
    difficulty,
    tags
) => {
    await writeFile(
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

/**
 * @param {string} problemDir
 * @param {string} problemName
 * @param {string} code
 */
const javascript = async (problemDir, problemName, code) => {
    mkdir(join(problemDir, 'javascript'));
    await writeFile(join(problemDir, 'javascript', 'index.js'), code);

    mkdir(join(problemDir, 'javascript', 'test'));
    await writeFile(
        join(problemDir, 'javascript', 'test', 'index.js'),
        [
            'const func = require(\'../index\');',
            'const { deepStrictEqual } = require(\'power-assert\');',
            '',
            'const testList = [];',
            '',
            `describe('${problemName}', () => {`,
            '    for (const [idx, test] of Object.entries(testList)) {',
            '        it(`TestCase[${idx}]`, () => {',
            '            deepStrictEqual(func(test));',
            '        });',
            '    }',
            '});',
        ].join('\n')
    );
};

/**
 * @param {string} problemDir
 * @param {string} _problemName
 * @param {string} code
 */
const golang = async (problemDir, _problemName, code) => {
    mkdir(join(problemDir, 'golang'));
    await writeFile(
        join(problemDir, 'golang', 'solution.go'),
        ['package golang', '', code].join('\n')
    );

    const funName = code.match(/func (?<funName>\w+)\s*\(.*/).groups.funName;
    const parseFunName = funName.charAt(0).toUpperCase() + funName.slice(1);

    await writeFile(
        join(problemDir, 'golang', 'solution_test.go'),
        [
            'package golang',
            '',
            'import (',
            '\t"reflect"',
            '\t"testing"',
            ')',
            '',
            'var tests = []struct {}{}',
            '',
            `func Test${parseFunName}(t *testing.T) {`,
            '\tfor idx, test := range tests {',
            `\t\tactual := ${funName}()`,
            '\t\tif !reflect.DeepEqual(actual, test.output) {',
            '\t\t\tt.Errorf("TestCase[%d]: ", idx)',
            '\t\t}',
            '\t}',
            '}',
        ].join('\n')
    );
};

main().catch((e) => console.log(e.stack));
