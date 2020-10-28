const { join, resolve } = require('path');
const { readdirSync, statSync, readFileSync, writeFileSync } = require('fs');

const PROJECT_ROOT = join(resolve(__dirname, '..'));
const ALG_DIR = join(PROJECT_ROOT, 'algorithms');

// call the main function on the last line
const main = () => {
    const algList = getAlgList();
    const algTable = algList
        .map(({ id, title, filename, difficulty, tags }) =>
            [
                id,
                `[${title}](algorithms/${filename})`,
                difficulty,
                tags.join(' '),
            ].join('|')
        )
        .map((item) => '|' + item + '|')
        .join('\n');
    const algCount = algList.reduce(
        (acc, { difficulty }) => {
            if (difficulty.includes('Easy')) {
                return { ...acc, easy: acc.easy + 1 };
            }
            if (difficulty.includes('Medium')) {
                return { ...acc, medium: acc.medium + 1 };
            }
            if (difficulty.includes('Hard')) {
                return { ...acc, hard: acc.hard + 1 };
            }

            throw new Error(`unknown difficulty(${difficulty})`);
        },
        { easy: 0, medium: 0, hard: 0 }
    );

    const result = [
        '# Profile',
        '',
        '## Algorithms',
        '',
        '### Problems Solved',
        '',
        `Total: ${algList.length}, Easy: ${algCount.easy}, Medium: ${algCount.medium}, Hard: ${algCount.hard}`,
        '',
        '### List',
        '',
        '| Id  | Title | Difficulty | Tags |',
        '| :-: |  :-:  |    :-:     | :-:  |',
        algTable,
        '',
    ];

    writeFileSync(join(PROJECT_ROOT, 'profile.md'), result.join('\n'));
};

const getAlgList = () =>
    readdirSync(ALG_DIR)
        .filter((filename) => statSync(join(ALG_DIR, filename)).isDirectory())
        .map((filename) => {
            const content = readFileSync(join(ALG_DIR, filename, 'README.md'))
                .toString()
                .split('\n');

            const algName = content[0].match(/##\s\[(?<algName>.*)\]/).groups.algName;

            const difficultyList = content
                .find((str) => str.startsWith('**Difficulty:**'))
                .match(/`(.*)`/g);
            if (difficultyList === null) {
                throw new Error('can not get difficulty for ' + filename);
            }
            const difficulty = difficultyList[0];

            let tags = content
                .find((str) => str.startsWith('**Tags:**'))
                .match(/`(.*)`/g);
            if (tags === null) {
                tags = [];
            }

            return {
                filename,
                difficulty,
                tags,

                id: parseInt(algName.match(/^(?<id>\d+)/).groups.id),
                title: algName.match(/\d+\.\s(?<title>.*)/).groups.title,
            };
        });

main();
