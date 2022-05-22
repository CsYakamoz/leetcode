import fetch from 'node-fetch';

const titleSlugRegex = /https:\/\/leetcode\.com\/problems\/(?<titleSlug>.*)\//;

/**
 * @param {string} problemLink
 */
export default async (problemLink, language) => {
    const titleSlug = problemLink.match(titleSlugRegex).groups.titleSlug;
    const body = {
        operationName: 'questionData',
        variables: { titleSlug },
        query:
            'query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      paidOnly\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    enableDebugger\n    envInfo\n    libraryUrl\n    adminUrl\n    __typename\n  }\n}\n',
    };

    const {
        data: { question },
    } = await fetch('https://leetcode.com/graphql', {
        headers: { 'content-type': 'application/json' },
        body: JSON.stringify(body),
        method: 'POST',
    }).then((res) => res.json());

    return {
        problemName: question.questionFrontendId + '. ' + question.title,
        description: question.content.replace(/\r\n/g, '\n'),
        tags: question.topicTags.map((item) => item.name),
        code: question.codeSnippets.find((item) => item.langSlug === language)
            .code,
        difficulty: question.difficulty,
    };
};
