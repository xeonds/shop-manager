<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>Customer Management</h1>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-text>Analysis</el-text>
          </template>
          <el-row>
            <el-col :span="6">
              <el-statistic title="All customers" :value="customerList.length" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="Ratio of men to women" :value="138">
                <template #suffix>/100</template>
              </el-statistic>
            </el-col>
            <el-col :span="6">
              <el-statistic title="Total Transactions" :value="172000" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="Feedback number" :value="562">
                <template #suffix>
                  <el-icon style="vertical-align: -0.125em">
                    <ChatLineRound />
                  </el-icon>
                </template>
              </el-statistic>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <el-text>All Customers</el-text>
              <el-button plain type="primary" @click="showCreateCustomer">Create New Customer</el-button>
            </div>
          </template>
          <el-table :data="customerList" style="width: 100%">
            <el-table-column type="selection"></el-table-column>
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="name" label="Name"></el-table-column>
            <el-table-column prop="email" label="Email"></el-table-column>
            <el-table-column prop="phone" label="Phone"></el-table-column>
            <el-table-column label="Operation">
              <template #default="{ row }">
                <el-button link type="primary" @click="handleEditCustomer(row)">Edit</el-button>
                <el-button link type="primary" @click="handleDeleteCustomer(row)">Delete</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog v-model="dialogVisible.create" title="Create Customer">
      <el-form :rules="rules" label-width="120px">
        <el-form-item label="Name" prop="name">
          <el-input v-model="customer.name"></el-input>
        </el-form-item>
        <el-form-item label="Email" prop="email">
          <el-input v-model="customer.email"></el-input>
        </el-form-item>
        <el-form-item label="Phone" prop="phone">
          <el-input v-model="customer.phone"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm(customer)">Save</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog v-model="dialogVisible.edit" title="Edit Customer">
      <el-form :rules="rules" label-width="120px">
        <el-form-item label="Name" prop="name">
          <el-input v-model="customer.name"></el-input>
        </el-form-item>
        <el-form-item label="Email" prop="email">
          <el-input v-model="customer.email"></el-input>
        </el-form-item>
        <el-form-item label="Phone" prop="phone">
          <el-input v-model="customer.phone"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm(customer)">Save</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import customerAPI from '../../api/customer'

export default {
  name: 'CustomerManagement',
  data() {
    return {
      customer: {
        name: '',
        email: '',
        phone: ''
      },
      rules: {
        name: [
          { required: true, message: 'Please input customer name', trigger: 'blur' }
        ],
        email: [
          { required: true, message: 'Please input customer email', trigger: 'blur' },
          { type: 'email', message: 'Please input correct email format', trigger: ['blur', 'change'] }
        ],
        phone: [
          { required: true, message: 'Please input customer phone', trigger: 'blur' },
          { pattern: /^[0-9]*$/, message: 'Please input correct phone format', trigger: ['blur', 'change'] }
        ]
      },
      customerList: [
        { id: 1, name: 'John Doe', email: 'johndoe@example.com', phone: '1234567890' },
        { id: 2, name: 'Jane Doe', email: 'janedoe@example.com', phone: '0987654321' }
      ],
      dialogVisible: {
        create: false,
        edit: false
      }
    }
  },
  methods: {
    async showCustomerList() {
      try {
        const response = await customerAPI.getCustomers()
        this.customerList = response.data
      } catch (error) {
        console.error(error)
        this.$message.error('Get customer list failed')
      }
    },
    showCreateCustomer() {
      this.dialogVisible.create = true
    },
    showEditCustomer() {
      this.dialogVisible.edit = true
    },
    async submitForm(formData) {
      try {
        await this.$refs[formName].validate()
        await customerAPI.createCustomer(this.customer)
        this.$message.success('Create customer success')
        this.resetForm(formName)
        await this.showCustomerList()
      } catch (error) {
        console.error(error)
        this.$message.error('Create customer failed')
      }
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    handleEditCustomer(customer) {
      this.customer = customer
      this.dialogVisible.edit = true
    },
    handleDeleteCustomer(customer) {
      this.$confirm(`Are you sure to delete ${customer.name}?`, 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(async () => {
        try {
          await customerAPI.deleteCustomer(customer.id)
          this.$message.success('Delete customer success')
          await this.showCustomerList()
        } catch (error) {
          console.error(error)
          this.$message.error('Delete customer failed')
        }
      }).catch(() => {
        this.$message.info('Delete customer canceled')
      })
    }
  },
  mounted() {
    this.showCustomerList()
  }
}
</script>

<style lang="less" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.customer-view {
  max-width: 1000px;
  margin: 0 auto;

  .el-row {
    margin-bottom: 2rem;
  }
}
</style>
