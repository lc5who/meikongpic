<template>
	<view class="top_bar">
		<uni-icons type="back" size="30" @click="goback"></uni-icons>
		常见问题
	</view>
	<view>
		<view class="border">1.提示“下载成功”，在相册里找不到怎么办?</view>
	<view class="small">
	&emsp;问题偶有发生，是宿主程序(非小程序)接口兼容性问
	题造成，图片被设置成了最早的日期。可打开相册，并翻
	到最后一页，尝试查找刚才已下载的图片。
	</view>
	<view class="border">2.提示“下载失败”怎么办?</view>
	<view class="small">
	&emsp;问题偶有发生，且多见于网络环境差的情况，请尝试
在WiFi环境或网络稳定的地方下载。
3.动态视频看不到下载按钮是怎么回事?
本小程序仅提供现今最为通用的视频格式动态壁纸,
因手机型号及版本多样化，对动态视频的支持不尽相同,
此问题偶有出现在oppo/vivo部分的机型，如遇此问题,
可尝试使用其他机型下载。
	</view>
	<view class="border">3.动态壁纸下载后如何在手机设置?</view>
	<view class="small">
	&emsp;基于小程序的权限框架原因，本小程序无法支持直接
设置手机动态壁纸，可下载三方软件进行设置!如Vmoon
	</view>
	<view class="border">4.下载的素材不是高清是怎么回事?</view>
	<view class="small">
		&emsp;所有素材由创作者上传，我们不做压缩和改变。普通
下载的素材非原图。高清下载的素材均为原图，如该原图
不是高清素材，请联系我们，我们将督促创作者重新.上
传。
		</view>
		<view class="border">5.为什么有时候加载较慢?</view>
		<view class="small">
				&emsp;在4G或WiFi网络环境下，正常的加载速度应是极快
的。但如果网络状态不佳，加. 上高清素材的尺寸较大，会
偶尔出现加载时间过长。
				</view>
				<view class="border">6.为什么要看广告?</view>
				<view class="small">
								&emsp;图片存储和下载都是需要消耗服务器资源的，广告是创
作者积极创作的动力来源，请理解和支持!
								</view>
								<view class="border">7.为什么每天下载次数有限?</view>
								<view class="small">
								为防止恶意刷量，保证良好的用户使用体验，当天下
								载次数均会设限，对此我们也将积极优化，感谢理解和支
								持!
																</view>
<view class="border">8.我想入驻壁纸平台成为创作者</view>
<view class="small">
	我们欢迎符合条件的小伙伴入驻，可联系我们(小程序
	菜单:我的-联系客服)私信申请。另外，入驻平台是完全免
	费的，需特别甄别，谨防上当!
</view>
	</view>
</template>

<script setup>
	let goback= ()=>{
		uni.navigateBack()
	}
	const aaa =111;
</script>

<style lang="scss">
	.border{
		font-weight: 900;
	}
	.small{
		font-size: 16upx;
		margin-bottom: 30upx;
	}
</style>
