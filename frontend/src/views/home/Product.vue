<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>Product Management</h1>
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
          <el-table :data="filteredCustomerList" style="width: 100%">
            <el-table-column type="selection"></el-table-column>
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="name" label="Name"></el-table-column>
            <el-table-column prop="sex" label="Sex"></el-table-column>
            <el-table-column prop="email" label="Email"></el-table-column>
            <el-table-column prop="phone" label="Phone"></el-table-column>
            <el-table-column label="Operation">
              <template #header>
                <el-input v-model="search" placeholder="Type to search" />
              </template>
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
      <el-form :model="customer" :rules="rules" label-width="120px" ref="customerNew">
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
          <el-button type="primary" @click="submitForm('customerNew')">Save</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog v-model="dialogVisible.edit" title="Edit Customer">
      <el-form :model="customer" :rules="rules" label-width="120px" ref="customerEdit">
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
          <el-button type="primary" @click="submitForm('customerEdit')">Save</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import customerAPI from '../../api/customer'

export default {
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
        { id: 1, name: 'John Doe', email: 'johndoe@example.com', phone: '1234567890', premier: true, sex: "man" },
        { id: 2, name: 'Jane Doe', email: 'janedoe@example.com', phone: '0987654321', premier: true, sex: "woman" }
      ],
      dialogVisible: {
        create: false,
        edit: false
      },
      search: '',
    }
  },
  computed: {
    filteredCustomerList() {
      return this.customerList.filter(customer => {
        return customer.name.toLowerCase().includes(this.search.toLowerCase())
          || customer.email.toLowerCase().includes(this.search.toLowerCase())
          || customer.phone.toLowerCase().includes(this.search.toLowerCase())
          || customer.id.toString().includes(this.search)
      })
    },
    men2Women() {
      return this.customerList.filter(customer => { customer.sex == 'man' }).length
        + '/' + this.customerList.filter(customer => { customer.sex == 'man' }).length
    },
    premierCustomers() {
      return this.customerList.filter(customer => { customer.premier == true }).length
    }
  },
  methods: {
    async showCustomerList() {
      try {
        const response = await customerAPI.getCustomers()
        this.customerList = response.data
      } catch (error) {
        this.$message.error('Get customer list failed')
      }
    },
    showCreateCustomer() {
      this.dialogVisible.create = true
    },
    showEditCustomer() {
      this.dialogVisible.edit = true
    },
    async submitForm(formName) {
      try {
        await this.$refs[formName].validate()
        await customerAPI.createCustomer(this.customer)
        this.$message.success('Create customer success')
        await this.showCustomerList()
      } catch (error) {
        this.$message.error('Create customer failed: ' + error || 'Unknown error')
      } finally {
        this.resetForm(formName)
        this.dialogVisible.create = false
        this.dialogVisible.edit = false
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
          this.$message.error('Delete customer failed: ' + error || 'Unknown error')
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

