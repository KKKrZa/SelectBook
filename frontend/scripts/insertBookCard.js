export { insertBookCard }

import { serverBaseUrl } from "./store.js";

// let isbnTimer // 查询 isbn 定时器，防止请求过于频繁

// 封装 书本信息 组件

function insertBookCard(parentNode, book) {
    let elem = document.createElement('div')
    elem.className = 'bookCard'
    let child_ul = document.createElement('ul')

    let bookTitle = document.createElement('li')
    let bookIsbn = document.createElement('li')
    let bookPressName = document.createElement('li')
    let courseName = document.createElement('li')
    let courseTeachers = document.createElement('li')
    let courseCollegeName = document.createElement('li')
    let detailBtn = document.createElement('button')

    bookTitle.textContent = `《${book['title']}》`
    bookIsbn.textContent = book['isbn']
    bookPressName.textContent = book['pressName']
    courseName.textContent = `课程: ${book['courseName']}`
    courseTeachers.textContent = `任课老师: ${book['teacherNames']}`
    courseCollegeName.textContent = `开设学院: ${book['courseCollegeName']}`
    detailBtn.textContent = '查看详细信息'
    detailBtn.saveIsbn = book['isbn']
    detailBtn.saveParentNode = child_ul
    detailBtn.onclick = (e) => {
        let isbn = e.target['saveIsbn']

        let requestOptions = {
            method: 'GET',
            // headers: myHeaders,
            redirect: 'follow'
        };
        fetch(`${serverBaseUrl}/queryISBN?isbn=${book['isbn']}`, requestOptions)
            .then(response => response.text())
            .then((res) => JSON.parse(res))
            .then((result) => {
                console.log(result)
                let res = result['data']['Books']
                if (res.length == 0) {
                    alert('请重试')
                }
                res = res[0]
                let imgUrl = res['img'].split('?')[0]
                let author = res['author']
                let language = res['language']

                let authorElem = document.createElement('li')
                let languageElem = document.createElement('li')
                let imgElem = document.createElement('img')

                authorElem.textContent = `作者: ${author}`
                languageElem.textContent = `语言: ${language}`
                imgElem.src = imgUrl
                imgElem.alt = '书本图片'

                child_ul.appendChild(authorElem)
                child_ul.appendChild(languageElem)
                child_ul.appendChild(imgElem)
            })
            .catch(error => console.log('error', error));

    }

    child_ul.appendChild(bookTitle)
    child_ul.appendChild(bookIsbn)
    child_ul.appendChild(bookPressName)
    child_ul.appendChild(courseName)
    child_ul.appendChild(courseTeachers)
    child_ul.appendChild(courseCollegeName)
    child_ul.appendChild(detailBtn)

    elem.appendChild(child_ul)
    parentNode.appendChild(elem)
}