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

Application requirements:

- go (v1.25.3) 
- sqlite3

Running the application

`docker build -t language-learner .`
`docker run -d -p 127.0.0.1:8080:8080 language-learner:latest`

Starting Screen Options:
![Quiz Starting Screen](images/quiz_start.png)

Quiz Question
![Quiz Question](images/quiz_question.png)

Finished Quiz Results Screen
![Quiz Results Screen](images/quiz_results.png)