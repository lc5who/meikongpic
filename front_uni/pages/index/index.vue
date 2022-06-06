<template>
	<view class="container">
		<view class="banner">

		</view>
		<view class="search">
			inpu
		</view>
		<view class="author">
			<view class="author_title">
				推荐创作者
			</view>

				<scroll-view scroll-x="true" scroll-y="false" enable-flex="true" @scroll="scroll" class="author_scroll" style="height: 150upx;">
					<view v-for="index in 15" :key="index" style="display: inline-block;width: 100upx;height: 100upx;">
						<image src="../../static/c2.png" mode=""></image>
					</view>
				</scroll-view>

			<!-- 			<scroll-view scroll-x="true"  scroll-y="false" class="author_scroll">
				<view class="author_avatar" v-for="index in 15" :key="index">
					<image src="../../static/c2.png" mode=""></image>
				</view>
			</scroll-view> -->
		</view>
		<view class="content">
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				href: 'https://uniapp.dcloud.io/component/README?id=uniui'
			}
		},
		methods: {
			scroll: function(e) {
				console.log(e)
				this.old.scrollTop = e.detail.scrollTop
			},
		}
	}
</script>

<style scoped="">
	.container {
		display: flex;
		flex-direction: column;
		width: 100vw;
		height: 100vh;
		background-color: coral;
	}

	.banner {
		height: 400upx;
		background-color: bisque;
	}

	.search {
		height: 200upx;
		background-color: aliceblue;
	}

	.author {
		height: 200upx;
		background-color: chartreuse;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-around;
		/* padding: 20upx; */
	}

	.author_avatar {
		display: inline-block;
		width: 80upx;
		height: 80upx;
	}

	.author_avatar image {
		width: 80upx;
		height: 80upx;
		border-radius: 50%;
		border: solid 5upx white;
	}

	.content {
		padding: 20upx;
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		gap: 10upx;
	}

	.content_image image {
		height: 200upx;
		width: 100%;

	}
	
</style>
