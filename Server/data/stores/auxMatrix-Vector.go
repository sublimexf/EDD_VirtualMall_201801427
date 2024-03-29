//Vector de departamentos con sus tiendas
//Estructura intermedia para poder ordenar el vector a row major

package data

import (
	"fmt"
	"strconv"
	"strings"
)

type AuxMatrix struct {
	Matrix []AuxVector
}

type AuxVector struct {
	Department string
	Vector     []NodeVector
}

type AllDepartments struct {
	Department []string
}

func NewMatrix() *AuxMatrix {
	return &AuxMatrix{}
}

func NewAuxVector() *AuxVector {
	return &AuxVector{"", nil}
}

func (mt *AuxMatrix) addToMatrix(dept string, node []NodeVector) {
	aux := AuxVector{dept, node}
	mt.Matrix = append(mt.Matrix, aux)
}

func (dpt *AllDepartments) AddDepartmentAll(dept string) {
	if dept != "" {
		saved := false
		for i := 0; i < len(dpt.Department); i++ {
			if strings.ToLower(dept) == strings.ToLower(dpt.Department[i]) {
				saved = true
				break
			}
		}
		if !saved {
			dpt.Department = append(dpt.Department, dept)
		}
	}
}

//Insertar informacion en el vector
// > Linealizado, tiendas lista doble
// >vector de departamentos existentes (sin repetir)
func (mt *AuxMatrix) SetDataMatrix(data Data, alldpt *AllDepartments ) {
	for i := 0; i < len(data.Data); i++ {
		mt.OrderIndex(&data, i) 
		mt.SetDepartmentMatrix(data.Data[i].Index, data.Data[i].Department, alldpt)	
	}
}

//Ordena la matriz segun el indice A-Z
func (mt *AuxMatrix) OrderIndex(data *Data, index int){
if index != len(data.Data)-1 {
		actualIndex := []byte(strings.ToLower(data.Data[index].Index))
		nextIndex := []byte(strings.ToLower(data.Data[index + 1].Index))
		if len(nextIndex) != 0 {
			if actualIndex[0] > nextIndex[0] {
				first := append(data.Data[:index], data.Data[index + 1], data.Data[index])		
				data.Data = append(first[:len(first)], data.Data[index+2:]...)
			}
		}		
	}
}

//Obtiene los departamentos departamanetos
func (mt *AuxMatrix) SetDepartmentMatrix(id string, department []DepartmentMatriz, alldpt *AllDepartments) {
	for i := 0; i < len(department); i++ {
		dpt := department[i]
		alldpt.AddDepartmentAll(dpt.Name)
		mt.SetStoresAux(id, dpt.Name, dpt.Store)
	}
}

//Crea e ingresa las tiendas al vector
func (mt *AuxMatrix) SetStoresAux(idVector string, dept string, storeinfo []StoreMatriz) {
	node1 := NewnodeVector()
	node2 := NewnodeVector()
	node3 := NewnodeVector()
	node4 := NewnodeVector()
	node5 := NewnodeVector()
	previousQual := 0
	id := dept + idVector

	node1.ID = id + "1"
	node1.Stores = NewStoresList()
	node2.ID = id + "2"
	node2.Stores = NewStoresList()
	node3.ID = id + "3"
	node3.Stores = NewStoresList()
	node4.ID = id + "4"
	node4.Stores = NewStoresList()
	node5.ID = id + "5"
	node5.Stores = NewStoresList()

	for i := 0; i < len(storeinfo); i++ {
		str := storeinfo[i]

		if previousQual != 0 && previousQual == str.Qualifi {

			switch str.Qualifi {
			case 1:
				node1.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept, str.Logo)
			case 2:
				node2.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept, str.Logo)
			case 3:
				node3.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept, str.Logo)
			case 4:
				node4.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept, str.Logo)
			case 5:
				node5.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept, str.Logo)
			default:
				fmt.Println("error: ubiacion " + idVector + dept + strconv.Itoa(str.Qualifi) + " tienda: " + str.Name)
				fmt.Println("No se encontre calificacion")
			}
		} else {
			i--
		}

		previousQual = str.Qualifi
	}

	node := []NodeVector{*node1, *node2, *node3, *node4, *node5}
	mt.addToMatrix(id, node)
}

//???????????????????????????????????????
// Vector -> Matrix ... -> .json
//???????????????????????????????????????

// func (mt *AuxMatrix) MapDepartments(vector Vector) {
// 	prevIndex := ""
// 	for i := 0; i < len(vector.Vector); i++ {
// 		auxIndex := []byte(vector.Vector[i].ID)
// 		index := string(auxIndex[len(auxIndex)-2])

// 		if prevIndex == index {
			

// 		}
	
// 		prevIndex = index
// 	}
// }

// func (mt *AuxMatrix) setAuxMatrix(dept string, ) {

// }