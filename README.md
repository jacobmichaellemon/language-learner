# Language Learner

A go application to help you learn vocabulary in different languages in a gamified way. The application allows to choose a native language
and a target language, how many words 10, 25, 50, 100 to quiz on, and the level of difficult base on word relevance.

Currently supports:
<table>
  <tr>
    <td>
      <ul>
        <li>"Catalan"</li>
        <li>"Czech"</li>
        <li>"Danish"</li>
        <li>"German"</li>
        <li>"English"</li>
        <li>"Spanish"</li>
        <li>"Finnish"</li>
        <li>"French"</li>
        <li>"Irish"</li>
        <li>"Italian"</li>
      </ul>
    </td>
    <td>
      <ul>
        <li>"Kurdish"</li>
        <li>"Lithuanian"</li>
        <li>"Malagasy"</li>
        <li>"Dutch"</li>
        <li>"Norwegian"</li>
        <li>"Polish"</li>
        <li>"Portuguese"</li>
        <li>"Swedish"</li>
        <li>"Turkish"</li>
      </ul>
    </td>
  </tr>
</table>

The application uses the datasets stored on: https://www.wikdict.com/page/download

- All data is extracted from Wiktionary by the DBnary project. 

## Motivation
- Wanted to create a simple and easy to use alternative to subscription based language learning applications
- Scalable difficulty allows user to build a foundational knowledge and slowly progress at their own pace
- Configurable to change the direction of learning and easily swap between langauges

### Application requirements:

- go (v1.25.3) 
- sqlite3

### Quick Start

`git clone https://github.com/jacobmichaellemon/language-learner`

`cd language-learner`

`docker build -t language-learner .`

### Usage

`docker run -d -p 127.0.0.1:8080:8080 language-learner:latest`

Options:
- Native Language: the language the words will be displayed in
- Target Lanagues: the language the words will be guessed in
- Number of Questions: how long the vocab quiz will be
- Difficulty: how relevant the words are in common discussions


![Quiz Starting Screen](https://github.com/jacobmichaellemon/language-learner/blob/main/images/quiz_start.PNG?raw=true)



![Quiz Question](https://github.com/jacobmichaellemon/language-learner/blob/main/images/quiz_question.PNG?raw=true)



![Quiz Results Screen](https://github.com/jacobmichaellemon/language-learner/blob/main/images/quiz_results.PNG?raw=true)


### Contributing 

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.