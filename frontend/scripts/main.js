import { insertBookCard } from "./insertBookCard.js"
import { serverBaseUrl } from "./store.js"

const queryMajors = document.querySelector('#grade-form')
const queryBtn = document.querySelector('#query')
const majorSelectElem = document.querySelector('#major-select')
const bookCardContainer = document.querySelector('#bookCardContainer')

queryMajors.addEventListener('submit', (e) => {
    console.log('查询专业')
    e.preventDefault()

    let grade = document.querySelector('#grade-select').value
    // console.log(grade)
    if (grade == null) {
        return
    }

    let requestOptions = {
        method: 'GET',
        // headers: myHeaders,
        redirect: 'follow'
    };

    fetch(`${serverBaseUrl}/listMajors`, requestOptions)
        .then(response => response.text())
        .then(result => JSON.parse(result))
        .then((result) => {
            // console.log(result)
            let majorList = result['data']
            majorList.sort((a, b) => {
                return a['collegeName'] === b['collegeName'] ? 0 : a['collegeName'] < b['collegeName'] ? -1 : 1;
            })
            console.log(majorList)
            if (majorList == null) {
                return
            }

            // 清空原列表
            let child = majorSelectElem.lastChild
            while (child) {
                majorSelectElem.removeChild(child)
                child = majorSelectElem.lastChild
            }

            for (let item of majorList) {
                // console.log(item)
                if (item['grade'] >= grade) { // 高年级包含低年级专业
                    // console.log(item['grade'])
                    // console.log(item['majorId'])
                    let majorOption = document.createElement('option')
                    majorOption.value = item['majorId']
                    majorOption.textContent = item['majorName']

                    majorSelectElem.appendChild(majorOption)
                }
            }
        })
        .catch(error => console.log('error', error));
})

queryBtn.addEventListener('click', (e) => {
    // e.preventDefault()

    let grade = document.querySelector('#grade-select').value
    let major = document.querySelector('#major-select').value

    console.log(grade, major)

    if (!grade || !major) {
        return
    }

    let urlencoded = new URLSearchParams();
    urlencoded.append("grade", grade);
    urlencoded.append("majorId", major);

    let child = bookCardContainer.lastChild
    // console.log(child)
    while (child) {
        bookCardContainer.removeChild(child)
        child = bookCardContainer.lastChild
    }

    let requestOptions = {
        method: 'POST',
        // headers: myHeaders,
        body: urlencoded,
        redirect: 'follow'
    };

    fetch(`${serverBaseUrl}/queryBookList`, requestOptions)
        .then(response => response.text())
        .then(result => JSON.parse(result))
        .then((result) => {
            let bookList = result['data']['list']
            if (bookList == null) {
                alert('No data found.')
            }
            for (let book of bookList) {
                console.log(book['title'])
                insertBookCard(bookCardContainer, book)
            }
        })
        .catch(error => console.log('error', error));
})

