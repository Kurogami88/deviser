package main

/* PARSE IN VALUE
1. Table Name (First char uppercase)
2. Table Name
3. Table Name (First char uppercase)
4. Field structure

5. Table Name (First char uppercase)
6. Table Name (First char uppercase)
7. Table Name
8. Table Name (First char uppercase)
9. Table Name (First char uppercase)
10. Table Name
11. Table Name
12. Table Name

13. Table Name (First char uppercase)
14. Table Name (First char uppercase)
15. Table Name (First char uppercase)
16. Table Name
17. Table Name (First char uppercase)
18. Table Name
19. Table Name
20. Table Name
21. Table Name
22. Table Name

23. Table Name (First char uppercase)
24. Table Name (First char uppercase)
25. Table Name (First char uppercase)
26. Table Name
27. Table Name (First char uppercase)
28. Table Name
29. Table Name
30. Table Name
31. Table Name
32. Table Name

33. Table Name (First char uppercase)
34. Table Name (First char uppercase)
35. Table Name (First char uppercase)
36. Table Name
37. Table Name (First char uppercase)
38. Table Name
39. Table Name
40. Table Name

41. Table Name (First char uppercase)
42. Table Name (First char uppercase)
43. Table Name
44. Table Name (First char uppercase)
45. Table Name (First char uppercase)
46. Table Name
47. Table Name
48. Table Name
49. Table Name
50. Table Name

51. Table Name (First char uppercase)
52. Table Name (First char uppercase)
53. Table Name
54. Table Name (First char uppercase)
55. Table Name (First char uppercase)
56. Table Name

57. Table Name (First char uppercase)
58. Table Name (First char uppercase)
59. Table Name
61. Table Name (First char uppercase)
62. Table Name (First char uppercase)
63. Table Name
64. Table Name
65. Table Name

66. Table Name (First char uppercase)
67. Table Name (First char uppercase)
68. Table Name
69. Table Name (First char uppercase)
70. Table Name (First char uppercase)
71. Table Name
72. Table Name
73. Table Name

74. Table Name (First char uppercase)
75. Table Name (First char uppercase)
76. Table Name (First char uppercase)
*/

var dbBase = `
/***
	Author: Leong Kai Khee (Kurogami)
	Date: 2020

	Generated by DeviserGO
***/

package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//%s is the DB schema of '%s' table
//Do not change the struct name, it has to be similar to the DB table name
type %s struct {%s
}

//DB%sCreate inserts a DB record
func DB%sCreate(%s %s) (%s, error) {
	// Error if key already exist
	err := db.Create(&%s).Error
	if err != nil {
		return %s, err
	}

	return %s, nil
}

//DB%sRetrieveUnscopeCondition retrieve DB records fulfilling your condition, including soft-deleted records
func DB%sRetrieveUnscopeCondition(conditions map[string]interface{}) ([]%s, error) {
	%ss := []%s{}

	if len(conditions) == 0 {
		err := db.Unscoped().Find(&%ss).Error
		if err != nil {
			return %ss, err
		}
	} else {
		err := db.Unscoped().Where(conditions).Find(&%ss).Error
		if err != nil {
			return %ss, err
		}
	}

	return %ss, nil
}

//DB%sRetrieveCondition retrieve DB records fulfilling your condition
func DB%sRetrieveCondition(conditions string) ([]%s, error) {
	%ss := []%s{}

	// conditions "user='me' and age = 20"
	if conditions == "" {
		err := db.Find(&%ss).Error
		if err != nil {
			return %ss, err
		}
	} else {
		err := db.Where(conditions).Find(&%ss).Error
		if err != nil {
			return %ss, err
		}
	}

	return %ss, nil
}

//DB%sRetrieve retrieve all DB records
func DB%sRetrieve() ([]%s, error) {
	%ss := []%s{}

	err := db.Find(&%ss).Error
	if err != nil {
		return %ss, err
	}

	return %ss, nil
}

//DB%sUpdate updates a DB record with matching primary key
func DB%sUpdate(%s %s) (%s, error) {
	// Will do nothing if key does not exist
	err := db.Model(&%s).Updates(&%s).Error
	if err != nil {
		return %s, err
	}

	db.First(&%s)
	return %s, nil
}

//DB%sUpdateCondition updates DB records fulfilling your condition
func DB%sUpdateCondition(%s %s, conditions string) error {
	// Will do nothing if key does not exist
	// conditions "user='me' and age = 20"
	err := db.Model(%s{}).Where(conditions).Updates(&%s).Error
	if err != nil {
		return err
	}

	return nil
}

//DB%sDelete soft-deletes a DB record with matching primary key
func DB%sDelete(%s %s) (%s, error) {
	// Will do nothing if key does not exist
	err := db.Delete(&%s).Error
	if err != nil {
		return %s, err
	}

	return %s, nil
}

//DB%sDeleteUnscope permanently deletes a DB record with matching primary key
func DB%sDeleteUnscope(%s %s) (%s, error) {
	// Will do nothing if key does not exist
	err := db.Unscoped().Delete(&%s).Error
	if err != nil {
		return %s, err
	}

	return %s, nil
}

//DB%sDeleteUnscopeCondition permanently deletes a DB record matching the condition
func DB%sDeleteUnscopeCondition(conditions string) error {
	// Will do nothing if key does not exist
	// conditions "user='me' and age = 20"
	err := db.Unscoped().Where(conditions).Delete(&%s{}).Error
	if err != nil {
		return err
	}

	return nil
}
`

var dbFixedFields = `
	CreatedAt *time.Time ` + "`gorm:\"default:current_timestamp\" json:\"createdAt\"`" + `
	UpdatedAt *time.Time ` + "`gorm:\"default:current_timestamp\" json:\"updatedAt\"`" + `
	DeletedAt *time.Time ` + "`json:\"deletedAt\"`"
var dbTableField = `
	%s %s ` + "`%s`"
var dbTableFieldJSON = `json:"%s"`
var dbTableFieldGorm = `gorm:"%s"`
var dbTableFieldGormType = `type:%s%s` //type:varchar(100)
var dbTableFieldGormKey = `primary_key`
var dbTableFieldGormAuto = `auto_increment:` //auto_increment:false
var dbTableFieldGormDefault = `default:`     //default:'galeone'
var dbTableFieldGormCT = `current_timestamp`
