import path from 'path/posix'
import { rollup } from 'rollup'
import { plugins } from './utils/config'
import crypto from 'crypto'
import fs from 'fs'

async function hash(filePath: string, algorithm: string, digest: crypto.BinaryToTextEncoding): Promise<string> {
	return new Promise(resolve => {
		const rs = fs.createReadStream(filePath);
		const hash = crypto.createHash(algorithm);
		let hex;
		rs.on('data', hash.update.bind(hash));
		rs.on('end', () => {
			resolve(hash.digest(digest))
		});
	});
}

const main = async() => {
  const p =  path.resolve(__dirname, '..', "..", '..', 'tasks', '搜狗拼音', 'scraper.ts')
  const h = await hash(p, "sha512", "hex")
  const bundler = await rollup({
    input: path.resolve(__dirname, '..', '..', '..', 'tasks', '搜狗拼音', 'scraper.ts'),
    plugins,
    treeshake: true,
  })
  
  const resl = await bundler.write({
    file: `sogou.${h.slice(0,10)}.js`,
    format: 'cjs'
  })
  console.log(resl)
}

main()