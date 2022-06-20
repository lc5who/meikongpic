<template>
	<view class="main">
	<view class="author_avatar_box">
		<view>
		<image src="../../static/logo2.jpg" mode="" class="author_avatar_image"></image>
		</view>
		<view>
			<text @click="apitest()" class="click_me_text">&nbsp;我的</text>
		</view>
	</view>
	<view class="login_box">
		<view>
			<image src="../../static/logo2.jpg" mode="" class="author_avatar_image2"></image>
		</view>
		<view>
			<text class="click_login_text">点击登录</text>
		</view>
	</view>
	<view class="" style="margin-top: 50upx;">
		<uni-list class="list-box">
			<uni-list-item title="成为创作者" clickable  @click="linkJoinAuthor()" class="list_box_list">
				<template v-slot:body>
					<view class="list_view">
					<view>
					<uni-icons type="auth-filled" size="30" class="icon"></uni-icons>
					</view>
					<view>
				<text>成为创作者</text>
				</view>
				</view>
				</template>
			</uni-list-item>
			<uni-list-item title="下载记录" link to="" @click="downrecord()" class="list_box_list">
				<template v-slot:body>
					<view class="list_view">
					<view>
					<uni-icons type="folder-add" size="30" class="icon"></uni-icons>
					</view>
					<view>
				<text>下载记录</text>
				</view>
				</view>
				</template>
			</uni-list-item>
			<uni-list-item title="常见问题" link to="" @click="linkProblem()" class="list_box_list">
				<template v-slot:body>
					<view class="list_view">
					<view>
					<uni-icons type="help" size="30" class="icon"></uni-icons>
					</view>
					<view>
				<text>常见问题</text>
				</view>
				</view>
				</template>
			</uni-list-item>
			<uni-list-item title="联系客服" link to="" @click="" class="list_box_list">
				<template v-slot:body>
					<view class="list_view">
					<view>
					<uni-icons type="phone" size="30" class="icon"></uni-icons>
					</view>
					<view>
				<text>联系客服</text>
				</view>
				</view>
				</template>
			</uni-list-item>
			<uni-list-item title="我要吐槽" link to="" @click="linkFeedback()" class="list_box_list">
				<template v-slot:body>
					<view class="list_view">
					<view>
					<uni-icons type="mail-open-filled" size="30" class="icon"></uni-icons>
					</view>
					<view>
				<text>我要吐槽</text>
				</view>
				</view>
				</template>
			</uni-list-item>
		</uni-list>
	</view>
	
	<view class="" style="margin-top: 50upx;">
		<uni-list>
			<uni-list-item title="分享给好友" link to="" @click="" class="list_box_list">
				<template v-slot:body>
					<view class="list_view">
					<view>
					<uni-icons type="staff-filled" size="30" class="icon"></uni-icons>
					</view>
					<view>
				<text>分享给好友</text>
				</view>
				</view>
				</template>
			</uni-list-item>
		</uni-list>
	</view>
	</view>
</template>

<script setup>
	const aaa =111;
	const linkFeedback=()=>{
		uni.navigateTo({
			url: '../feedback/feedback'
		});
	}
	const linkProblem=()=>{
		uni.navigateTo({
			url: './problem',
		});
	}
	const linkJoinAuthor=()=>{
		uni.navigateTo({
			url: '/pages/author/joinauthor/joinauthor'
		});
	}
	const downrecord=()=>{
		uni.navigateTo({
			url: '/pages/me/downrecord/downrecord'
		});
	}
	const apitest=()=>{
		uni.request({
		  url: "http://101.36.111.54:1031/api/test",
		  success: (res) => {
		    console.log(res);
		  },
		  fail: (err) => {
		    console.log(err);
		  },
		});
	}
</script>

<style lang="scss">
	body{
		background-color: rgb(20, 20, 58);
	}
	.main{
		background-color: rgb(20, 20, 58);
	}
.author_avatar_box {
		width: 160upx;
		height: 70upx;
		display: flex;
		align-items: center;
		justify-content: flex-start;
	}
	.author_avatar_image {
		width: 50upx;
		height: 50upx;
		border-radius: 50%;
		background-color: white;
	}
	.author_avatar_image2 {
		width: 130upx;
		height: 130upx;
		border-radius: 50%;
		background-color: white;
	}
	.login_box{
		width: 700upx;
		height: 150upx;
		background-color: rgb(47, 38, 155);
		border-radius: 20upx;
		margin-left: 25upx;
		margin-top: 30upx;
		display: flex;
		align-items: center;
		justify-content: flex-start;
	}
	.click_login_text{
		font-size: 40upx;
		margin-left: 50upx;
		color: white;
	}
	.click_me_text{
		color: white;
	}
	.list-box {
			background-color: black;
		}
	.list_box_list{
		background-color: rgb(27, 27, 27);
		color: white;
		// border-radius: 10px;
	}
	.list_view {
			display: flex;
			align-items: center;
			justify-content: flex-start;
		}
		.icon{
			margin-right: 40upx;
			color: white;
		}
</style>
