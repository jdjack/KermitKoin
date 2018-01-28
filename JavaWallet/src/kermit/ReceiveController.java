package kermit;

import javafx.fxml.FXML;
import javafx.scene.text.Text;

public class ReceiveController {

  String address;
  @FXML
  private Text addressField;

  public void setAddress(String address) {
    this.address = address;
    addressField.setText(address);

  }
}
