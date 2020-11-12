package initialize

import (
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

// 初始化数据
func InitData() {
	// 1. 初始化角色
	creator := "系统创建"
	status := true
	roles := []models.SysRole{
		{
			Model: models.Model{
				Id: 1,
			},
			Name:    "管理员",
			Keyword: "admin",
			Desc:    "系统管理员",
			Status:  &status,
			Creator: creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:    "访客",
			Keyword: "guest",
			Desc:    "外来访问人员",
			Status:  &status,
			Creator: creator,
		},
	}
	for _, role := range roles {
		oldRole := models.SysRole{}
		err := common.Mysql.Where("id = ?", role.Id).First(&oldRole).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&role)
		}
	}

	// 2. 初始化菜单
	menus := []models.SysMenu{
		{
			Model: models.Model{
				Id: 1,
			},
			Name:     "仪表盘",
			Icon:     "icon-yibiaopan",
			Path:     "index",
			Sort:     0,
			Status:   &status,
			ParentId: 0,
			Creator:  creator,
			Roles:    roles,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:     "系统设置",
			Icon:     "icon-icon_shezhi",
			Path:     "system",
			Sort:     1,
			Status:   &status,
			ParentId: 0,
			Creator:  creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
		{
			Model: models.Model{
				Id: 3,
			},
			Name:     "用户管理",
			Icon:     "icon-icon_zhanghao",
			Path:     "user",
			Sort:     10,
			Status:   &status,
			ParentId: 2,
			Creator:  creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
		{
			Model: models.Model{
				Id: 4,
			},
			Name:     "部门管理",
			Icon:     "icon-icon_zhanghao",
			Path:     "dept",
			Sort:     11,
			Status:   &status,
			ParentId: 2,
			Creator:  creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
		{
			Model: models.Model{
				Id: 5,
			},
			Name:     "菜单管理",
			Icon:     "icon-wuxupailie",
			Path:     "menu", // 子菜单不用全路径, 自动继承
			Sort:     12,
			Status:   &status,
			ParentId: 2,
			Creator:  creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
		{
			Model: models.Model{
				Id: 6,
			},
			Name:     "角色管理",
			Icon:     "icon-quanxianshenpi",
			Path:     "role",
			Sort:     13,
			Status:   &status,
			ParentId: 2,
			Creator:  creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
		{
			Model: models.Model{
				Id: 7,
			},
			Name:     "接口管理",
			Icon:     "icon-APIshuchu",
			Path:     "api",
			Sort:     14,
			Status:   &status,
			ParentId: 2,
			Creator:  creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
		{
			Model: models.Model{
				Id: 8,
			},
			Name:     "操作日志",
			Icon:     "icon-gaojing1",
			Path:     "log",
			Sort:     15,
			Status:   &status,
			ParentId: 2,
			Creator:  creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
	}
	for _, menu := range menus {
		oldMenu := models.SysMenu{}
		err := common.Mysql.Where("id = ?", menu.Id).First(&oldMenu).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&menu)
		}
	}

	// 3. 初始化用户
	// 默认头像
	avatar := "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	users := []models.SysUser{
		{
			Username: "admin",
			Password: utils.GenPwd("123456"),
			Mobile:   "18888888888",
			Avatar:   avatar,
			Name:     "管理员",
			RoleId:   1,
			Creator:  creator,
		},
		{
			Username: "guest",
			Password: utils.GenPwd("123456"),
			Mobile:   "15888888888",
			Avatar:   avatar,
			Name:     "来宾",
			RoleId:   2,
			Creator:  creator,
		},
	}
	for _, user := range users {
		oldUser := models.SysUser{}
		err := common.Mysql.Where("username = ?", user.Username).First(&oldUser).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&user)
		}
	}

	// 4. 初始化接口
	apis := []models.SysApi{
		{
			Model: models.Model{
				Id: 1,
			},
			Method:     "POST",
			Path:       "/auth/login",
			Category:   "auth",
			Desc:       "获取用户信息和token",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Method:     "POST",
			Path:       "/auth/logout",
			Category:   "auth",
			Desc:       "用户登出",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 3,
			},
			Method:     "POST",
			Path:       "/auth/refresh_token",
			Category:   "auth",
			Desc:       "刷新JWT令牌",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 4,
			},
			Method:     "POST",
			Path:       "/v1/user/info",
			Category:   "user",
			Desc:       "用户信息",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 5,
			},
			Method:     "PATCH",
			Path:       "/v1/user/info/update/:userId",
			Category:   "user",
			Desc:       "更新信息",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 6,
			},
			Method:     "POST",
			Path:       "/v1/user/info/uploadImg",
			Category:   "user",
			Desc:       "上传头像",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 7,
			},
			Method:     "GET",
			Path:       "/v1/user/list",
			Category:   "user",
			Desc:       "用户列表",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 8,
			},
			Method:     "PUT",
			Path:       "/v1/user/changePwd",
			Category:   "user",
			Desc:       "修改密码",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 9,
			},
			Method:     "POST",
			Path:       "/v1/user/create",
			Category:   "user",
			Desc:       "创建用户",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 10,
			},
			Method:     "PATCH",
			Path:       "/v1/user/update/:userId",
			Category:   "user",
			Desc:       "更新用户",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 11,
			},
			Method:     "DELETE",
			Path:       "/v1/user/delete",
			Category:   "user",
			Desc:       "删除用户",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 12,
			},
			Method:     "GET",
			Path:       "/v1/menu/tree",
			Category:   "menu",
			Desc:       "获取菜单",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 13,
			},
			Method:     "GET",
			Path:       "/v1/menu/list",
			Category:   "menu",
			Desc:       "菜单列表",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 14,
			},
			Method:     "POST",
			Path:       "/v1/menu/create",
			Category:   "menu",
			Desc:       "创建菜单",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 15,
			},
			Method:     "PATCH",
			Path:       "/v1/menu/update/:menuId",
			Category:   "menu",
			Desc:       "更新菜单",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 16,
			},
			Method:     "DELETE",
			Path:       "/v1/menu/delete",
			Category:   "menu",
			Desc:       "删除菜单",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 17,
			},
			Method:     "GET",
			Path:       "/v1/role/list",
			Category:   "role",
			Desc:       "角色列表",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 18,
			},
			Method:     "POST",
			Path:       "/v1/role/create",
			Category:   "role",
			Desc:       "创建角色",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 19,
			},
			Method:     "PATCH",
			Path:       "/v1/role/update/:roleId",
			Category:   "role",
			Desc:       "更新角色",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 20,
			},
			Method:     "PATCH",
			Path:       "/v1/role/perms/update/:roleId",
			Category:   "role",
			Desc:       "更新权限",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 21,
			},
			Method:     "DELETE",
			Path:       "/v1/role/delete",
			Category:   "role",
			Desc:       "删除角色",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 22,
			},
			Method:     "GET",
			Path:       "/v1/api/list",
			Category:   "api",
			Desc:       "获取接口",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 23,
			},
			Method:     "POST",
			Path:       "/v1/api/create",
			Category:   "api",
			Desc:       "创建接口",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 24,
			},
			Method:     "PATCH",
			Path:       "/v1/api/update/:apiId",
			Category:   "api",
			Desc:       "更新接口",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 25,
			},
			Method:     "DELETE",
			Path:       "/v1/api/delete",
			Category:   "api",
			Desc:       "删除接口",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 26,
			},
			Method:     "PATCH",
			Path:       "/v1/role/menus/update/:roleId",
			Category:   "role",
			Desc:       "更新菜单",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 27,
			},
			Method:     "PATCH",
			Path:       "/v1/role/apis/update/:roleId",
			Category:   "role",
			Desc:       "更新权限",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 28,
			},
			Method:     "GET",
			Path:       "/v1/operlog/list",
			Category:   "operLog",
			Desc:       "获取操作日志",
			Creator:    creator,
		},
		{
			Model: models.Model{
				Id: 29,
			},
			Method:     "DELETE",
			Path:       "/v1/operlog/delete",
			Category:   "operLog",
			Desc:       "删除操作日志",
			Creator:    creator,
		},
	}
	for _, api := range apis {
		oldApi := models.SysApi{}
		err := common.Mysql.Where("id = ?", api.Id).First(&oldApi).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&api)
			s := service.New()
			// 管理员拥有所有API权限role[0]
			_, err = s.CreateRoleCasbin(models.SysRoleCasbin{
				Keyword: roles[0].Keyword,
				Path:    api.Path,
				Method:  api.Method,
			})
			// 来宾权限
			if api.Id <= 6 || api.Id == 9 {
				_, err = s.CreateRoleCasbin(models.SysRoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
			}
		}
	}
}
