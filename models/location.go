package models

import "time"

// Location is a GPS point on a map.
type Location struct {
	Model
	// User      User      `json:"user"`
	UserID    int       `gorm:"one2many:User;AssociationForeignKey:id;ForeignKey:user_id" json:"user_id"`
	Location  GeoPoint  `json:"location" sql:"type:geometry(Geometry,4326)"`
	Timestamp time.Time `json:"timestamp"`
}

// GeoPoint maps against Postgis geographical point
// Note: Longitude (x) is stored in Postgres before Latitude (y).
// type GeoPoint struct {
// 	Lat float64 `json:"lat"`
// 	Lng float64 `json:"lng"`
// }
//
// func (p *GeoPoint) String() string {
// 	log.Printf("[DEBUG] creating point")
// 	log.Printf("[DEBUG] POINT(%v, %v) \n", p.Lat, p.Lng)
// 	return fmt.Sprintf(`SRID=4326;POINT(%v, %v)`, p.Lat, p.Lng)
// }
//
// // Scan implements the Scanner interface and will scan the Postgis POINT(x y) into the GeoPoint struct
// func (p *GeoPoint) Scan(val interface{}) error {
// 	rawPoint := val.([]byte)
// 	sep := bytes.IndexByte(rawPoint, ',')
// 	var err error
// 	if p.Lat, err = strconv.ParseFloat(string(rawPoint[1:sep]), 64); err != nil {
// 		return err
// 	}
//
// 	if p.Lng, err = strconv.ParseFloat(string(rawPoint[sep+1:len(rawPoint)-1]), 64); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // Value implements the driver Valuer interface and will return the string representation of the GeoPoint struct by calling the String() method
// func (p GeoPoint) Value() (driver.Value, error) {
// 	log.Printf("[DEBUG] creating point")
// 	return p.String(), nil
// }
