import EventEmitter from 'events'
import chalk from 'chalk'
export const __logEvent = new EventEmitter()

export enum LogLevel {
  Info, Success, Warning,
  Error, Normal
}

export function log(text: string) {
	// 增加字符串类型判断
	if (typeof text !== 'string') {
		__logEvent.emit("log", LogLevel.Warning, 'Illegal type detected');
		__logEvent.emit("log", LogLevel.Normal, JSON.stringify(text));
		return;
	}

	const spl = text.split(':');
	if (spl.length < 2) {
		__logEvent.emit("log", LogLevel.Warning, 'Illegal message detected');
    __logEvent.emit("log", LogLevel.Normal, text);
		return;
	}

	const inf = text.substring(spl[0].length + 1);
	switch (spl[0]) {
		case 'Info':
			// if (args.hasOwnProperty('g')) {
				__logEvent.emit("log", LogLevel.Info, inf);
			// } else {
			// 	console.log(chalk.blue('Info ') + inf);
			// }

			break;
		case 'Success':
			// if (args.hasOwnProperty('g')) {
				__logEvent.emit("log", LogLevel.Success, + inf);
			// } else {
			// 	console.log(chalk.greenBright('Success ') + inf);
			// }

			break;
		case 'Warning':
			// if (args.hasOwnProperty('g')) {
			// 	console.log('::warning::' + inf);
			// } else {
				__logEvent.emit("log", LogLevel.Warning, inf);
			// }

			break;
		case 'Error':
			// if (args.hasOwnProperty('g')) {
			// 	console.log('::error::' + inf);
			// } else {
				__logEvent.emit("log", LogLevel.Error, inf);
			// }

			break;
		default:
			// if (args.hasOwnProperty('g')) {
			// 	console.log('::warning::Illegal message detected:' + inf);
			// } else {
				__logEvent.emit("log", LogLevel.Warning, 'Illegal message detected');
				__logEvent.emit("log", LogLevel.Normal, text)
			// }
	}
}

import fs from 'fs';
import cpt from 'crypto';
import iconv from 'iconv-lite';
import {Cmp, Status} from './enum';
import path from "path";
import {Interface} from './class'

async function getMD5(filePath: string): Promise<string> {
	return new Promise(resolve => {
		const rs = fs.createReadStream(filePath);
		const hash = cpt.createHash('md5');
		let hex;
		rs.on('data', hash.update.bind(hash));
		rs.on('end', () => {
			hex = hash.digest('hex');
			log('Info:MD5 is ' + hex);
			resolve(hex);
		});
	});
}

function parseDownloadUrl(href: string): string {
	// 识别根目录字符“/”
	if (href[0] === '/') {
		href = 'https://portableapps.com' + href;
	}

	// 识别downloading，替换为redirect
	href = href.replace('downloading', 'redirect');

	// URL编码
	href = encodeURI(href);

	// Log("Info:Parse download link into:" + href)
	return href;
}

function formatVersion(version: string): string {
	const spl = version.split('.');
	if (spl.length > 4) {
		log('Warning:Illegal version "' + version + ',"length=' + spl.length);
		return version;
	}

	// 将版本号扩充为4位
	for (let i = 0; i < 4 - spl.length; i++) {
		version += '.0';
	}

	return version;
}

function matchVersion(text: string): Interface {
	const regex = /\d+.\d+(.\d+)*/;
	const matchRes = text.match(regex);
	if (!matchRes || matchRes.length === 0) {
		return new Interface({
			status: Status.ERROR,
			payload:
				'Error:Matched nothing when looking into "'
				+ text
				+ '" with "'
				+ regex
				+ '",skipping...',
		});
	}

	return new Interface({
		status: Status.SUCCESS,
		payload: matchRes[0],
	});
} // Interface:string

function isURL(str_url: string): boolean {
	return str_url.slice(0, 4) == "http"
}

function getSizeString(size: number): string {
	if (size < 1024) return size.toFixed(2) + "B"
	else if (size < 1024 * 1024) return (size / 1024).toFixed(2) + "KB"
	else if (size < 1024 * 1024 * 1024) return (size / (1024 * 1024)).toFixed(2) + "MB"
	else return (size / (1024 * 1024 * 1024)).toFixed(2) + "GB"
}

function versionCmp(a: string, b: string): Cmp {
	const x = a.split('.');
	const y = b.split('.');
	let result: Cmp = Cmp.E;

	for (let i = 0; i < Math.min(x.length, y.length); i++) {
		if (Number(x[i]) < Number(y[i])) {
			result = Cmp.L;
			break;
		} else if (Number(x[i]) > Number(y[i])) {
			result = Cmp.G;
			break;
		}
	}

	// 处理前几位版本号相同但是位数不一致的情况，如1.3/1.3.0
	if (result === Cmp.E && x.length !== y.length) {
		// 找出较长的那一个
		let t: Array<string>;
		t = x.length < y.length ? y : x;
		// 读取剩余位
		for (
			let i = Math.min(x.length, y.length);
			i < Math.max(x.length, y.length);
			i++
		) {
			if (Number(t[i]) !== 0) {
				result = x.length < y.length ? Cmp.L : Cmp.G;
				break;
			}
		}
	}

	return result;
}

function rd(dst: string): boolean {
	throw new Error('not implemented');
}

function mv(src: string, dst: string): boolean {
	throw new Error('not implemented');
}

function xcopy(src: string, dst: string): boolean {
	throw new Error('not implemented');
}

function cleanBuildStatus(s: any[]): any[] {
	throw new Error('not implemented');
}

function gbk(buffer: Buffer): string {
	return iconv.decode(buffer, 'GBK');
}

function gb2312(buffer: Buffer): string {
	return iconv.decode(buffer, 'GB2312');
}

function toGbk(text: string): Buffer {
	return iconv.encode(text, 'GBK');
}

function copyCover(task: any, p7zip: string): boolean {
	throw new Error('not implemented');
}

async function awaitWithTimeout(closure: () => any, timeout: number): Promise<any> {
	return new Promise((async (resolve, reject) => {
		setTimeout(() => {
			reject("Error:External scraper init() timeout")
		}, timeout)
		let res = await closure()
		resolve(res)
	}))
}

function printMS(ms: number): string {
	const s = ms / 1000
	if (s < 60) {
		return `${s.toFixed(1)} s`
	} else {
		return `${(s / 60).toFixed(1)} min`
	}
}

export {
	getMD5,
	parseDownloadUrl,
	formatVersion,
	matchVersion,
	versionCmp,
	rd,
	mv,
	xcopy,
	cleanBuildStatus,
	gbk,
	gb2312,
	toGbk,
	copyCover,
	isURL,
	getSizeString,
	awaitWithTimeout,
	printMS
};
