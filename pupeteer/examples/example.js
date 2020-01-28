const puppeteer = require('puppeteer');

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
