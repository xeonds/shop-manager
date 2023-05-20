<template>
	<el-container class="container">
		<el-card class="login-card">
			<template #header>
				<el-text>登录</el-text>
			</template>
			<el-form :rules="rules">
				<el-form-item label="用户名">
					<el-input placeholder="请输入用户名" v-model="form.username"></el-input>
				</el-form-item>
				<el-form-item label="密码">
					<el-input placeholder="请输入密码" v-model="form.password"></el-input>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" plain @click="onSubmit">登录</el-button>
				</el-form-item>
			</el-form>
		</el-card>
	</el-container>
</template>

<script>
import admin from '../api/admin';

export default {
	data() {
		return {
			form: {
				username: '',
				password: ''
			},
			rules: {
				username: [
					{ required: true, message: '请输入用户名', trigger: 'blur' },
					{ min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
				],
				password: [
					{ required: true, message: '请输入密码', trigger: 'blur' },
					{ min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
				]
			}
		};
	},
	methods: {
		async onSubmit() {
			await admin.login(formData.username, formData.password)
				.then(res => {
					if (res.code == 200) {
						this.$message({
							message: '登录成功',
							type: 'success'
						})
						this.$router.push({ path: '/main/dashboard' })
					} else {
						this.$message({
							message: res.message,
							type: 'error'
						})
					}
				})
				.catch(err => {
					console.log(err)
				})
		}
	},
};
</script>

<style lang="less" scoped>
.el-container {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100vh;
}
</style>
