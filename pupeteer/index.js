const puppeteer = require('puppeteer');

(async () => {
  const browser = await puppeteer.launch(
    { headless: true }
  )
  const url = process.env.URL
  const page = await browser.newPage()
  await page.goto(url)
  const items = await page.$$('tr.default')
  const data = await Promise.all(
    items.map(
      async (item) => {
        const itemColumns = await item.$$('td')
        const categoryElement = await itemColumns[0].$('a')
        const categoryLink = await (await categoryElement.getProperty('href')).jsonValue()
        const category = new URL(categoryLink).search.replace(/\?c\=/, '')
        const nameElement = await itemColumns[1].$('td > a')
        const nameLink = await (await nameElement.getProperty('href')).jsonValue()
        const nameText = await (await nameElement.getProperty('title')).jsonValue()

        const sizeElement = itemColumns[3]
        // TODO: parse size
        const size = await (await sizeElement.getProperty('textContent')).jsonValue()

        const dateElement = itemColumns[4]
        // TODO: parse dateTime
        const createdAt = await (await dateElement.getProperty('textContent')).jsonValue()

        const seedersElement = itemColumns[5]
        const seeders = await (await seedersElement.getProperty('textContent')).jsonValue()
        const leechersElement = itemColumns[6]
        const leechers = await (await leechersElement.getProperty('textContent')).jsonValue()
        const completedDownloadsElement = itemColumns[7]
        const completedDownloads = await (await completedDownloadsElement.getProperty('textContent')).jsonValue()
        return {
          category,
          nameLink,
          nameText,
          size,
          createdAt,
          seeders,
          leechers,
          completedDownloads
        }
      }
    )
  )

  console.log(data)

  await page.close()
  await browser.close()
})();
/**
(async () => {
  const browser = await puppeteer.launch({ headless: false });
  const page = await browser.newPage()

  // await page.goto('https://news.ycombinator.com/news')
  await page.goto('http://example.com')

  // const data = await page.$eval('.hnname > a', el => el.innerText)
  // const item = await page.$('p')
  // const data = await (await item.getProperty('textContent')).jsonValue()

  // const item = await page.$('p')
  // const data = await (await item.getProperty('textContent')).jsonValue()

  const items = await page.$$('p')
  // const data = items.map((item) => {
  //   return await (await item.getProperty('textContent')).jsonValue()
  // })
  const data = await Promise.all(
    items.map(async (item) => {
      const contentRaw = await item.getProperty('textContent')
      return await contentRaw.jsonValue()
    })
  )

  // const item = await page.$('table')
  // const data = await (await item.getProperty('textContent')).jsonValue()
  // const item = await page.$('table')
  // const data = await (await item.getProperty('textContent')).jsonValue()

  console.log(data)
  console.log(data.length)
  await page.waitFor(100000)
  await browser.close()
})();
*/
